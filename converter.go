package apptree

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func (r *Record) WriteInto(v interface{}) error {
	return WriteInto(v, r)
}

func (r *Record) ReadFrom(v interface{}) error {
	return ReadFrom(v, r)
}

func ReadFrom(v interface{}, rec *Record) error {
	structValue := reflect.ValueOf(v)
	structDesc := reflect.TypeOf(v)
	switch structValue.Kind() {
	case reflect.Ptr:
		structValue = structValue.Elem()
		structDesc = structDesc.Elem()
	}
	for i := 0; i < structDesc.NumField(); i++ {
		structField := structDesc.Field(i)
		fieldValue := structValue.Field(i)
		tag := structField.Tag
		dataSetItemIndex, skip, err := getIndexFromTag(tag)
		if err != nil {
			return err
		}
		if skip {
			continue
		}
		if typedVal, ok := fieldValue.Interface().(TypedValue); ok {
			rec.setValue(typedVal, dataSetItemIndex)
			continue
		}
		switch fieldValue.Kind() {
		case reflect.Slice:
			configAttr := rec.Configuration.getConfigurationAttribute(dataSetItemIndex)
			if configAttr == nil {
				return errors.New("No attribute defined")
			}
			if configAttr.Type != Type_Relationship {
				return fmt.Errorf("Attribute %d is a slice but is a type %s. Expected %s", dataSetItemIndex, configAttr.Type, Type_Relationship)
			}
			for childIndex := 0; childIndex < fieldValue.Len(); childIndex++ {
				childRec, err := rec.AddToManyChildAtIndex(dataSetItemIndex)
				if err != nil {
					return err
				}
				sliceRec := fieldValue.Index(childIndex)
				ReadFrom(sliceRec.Interface(), childRec)
			}
		case reflect.Struct:
			configAttr := rec.Configuration.getConfigurationAttribute(dataSetItemIndex)
			if configAttr == nil {
				return errors.New("No attribute defined")
			}
			if configAttr.Type != Type_SingleRelationship {
				return fmt.Errorf("Attribute %d is a struct but is a type %s. Expected %s", dataSetItemIndex, configAttr.Type, Type_SingleRelationship)
			}
			childRec, err := rec.NewToOneRelationshipAtIndex(dataSetItemIndex)
			if err != nil {
				return err
			}
			ReadFrom(fieldValue.Interface(), childRec)
		}
	}
	return nil
}

func WriteInto(v interface{}, rec *Record) error {
	structValue := reflect.ValueOf(v)
	structDesc := reflect.TypeOf(v)
	switch structValue.Kind() {
	case reflect.Ptr:
		structValue = structValue.Elem()
		structDesc = structDesc.Elem()
	}

	for i := 0; i < structValue.NumField(); i++ {
		structField := structDesc.Field(i)
		fieldValue := structValue.Field(i)
		tag := structField.Tag
		dataSetItemIndex, skip, err := getIndexFromTag(tag)
		if err != nil {
			return err
		}
		if skip {
			continue
		}
		typedValue := rec.getValue(dataSetItemIndex)
		if relation, ok := typedValue.(ToManyRelationship); ok {
			if fieldValue.Kind() != reflect.Slice {
				return fmt.Errorf("Attribute %i of %s is of type ToManyRelationship but the Struct attribute is not a slice", i, rec.Configuration.Name)
			}
			destinationType := fieldValue.Type().Elem()
			for _, item := range relation.Items {
				childStruct := reflect.New(destinationType)
				WriteInto(childStruct.Interface(), &item)
				fieldValue.Set(reflect.Append(fieldValue, childStruct.Elem()))
			}
		} else if relation, ok := typedValue.(SingleRelationship); ok {
			destinationType := fieldValue.Type()
			childStruct := reflect.New(destinationType)
			WriteInto(childStruct.Interface(), &relation.Record)
			fieldValue.Set(childStruct.Elem())
		} else {
			recValue := typedValue.(interface{})
			reflectValue, ok := recValue.(reflect.Value)
			if !ok {
				reflectValue = reflect.ValueOf(recValue)
			}
			if reflectValue.IsValid() {
				if reflectValue.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(reflectValue.Convert(fieldValue.Type()))
				} else {
					if fieldValue.Kind() == reflect.Ptr {
						if fieldValue.IsNil() {
							fieldValue.Set(reflect.New(structField.Type))
						}
						fieldValue = fieldValue.Elem()
					}

					if reflectValue.Type().ConvertibleTo(fieldValue.Type()) {
						fieldValue.Set(reflectValue.Convert(fieldValue.Type()))
					} else {
						err = fmt.Errorf("could not convert argument of field %s from %s to %s", structField.Name, reflectValue.Type(), fieldValue.Type())
					}
				}
			}
		}
	}
	return nil
}

func getIndexFromTag(tag reflect.StructTag) (int, bool, error) {
	indexStr := tag.Get("index")
	if indexStr == "" || indexStr == "-" {
		return 0, true, nil
	}
	index64, err := strconv.ParseInt(indexStr, 10, 32)
	if err != nil {
		return 0, false, err
	}
	return int(index64), false, nil
}
