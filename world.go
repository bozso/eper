package eper

import (
    "reflect"
)

type World struct {
    TypeMap
    Components
    typeCounter TypeID
}

func NewWorld() (w World) {
    return World {
        TypeMap: make(TypeMap),
        Components: make(Components),
        typeCounter: 0,
    }
}

func (w *World) addType(t reflect.Type) (id TypeID) {
    id = w.typeCounter
    w.TypeMap[t] = id
    w.typeCounter++

    return id
}

func (w *World) AddType(e Entity) (id TypeID) {
    return w.addType(reflect.TypeOf(e))
}

func (w *World) Set(id ID, e Entity) (err error) {
    return
}
