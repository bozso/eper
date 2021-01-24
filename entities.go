package eper

import (
    "reflect"
)

type TypeID ID
type TypeMap map[reflect.Type]TypeID

type Entity struct {
    generation uint32
    id uint32
}

func (e Entity) ToBits() uint64 {
    return uint64(e.generation) << 32 | uint64(e.id)
}

func EntityFromBits(bits uint64) (e Entity) {
    return Entity {
        generation: uint32(bits >> 32),
        id: uint32(bits),
    }
}

type Entities map[ID]Entity


type Getter interface {
    Get(ID, Entity) error
}

type Setter interface {
    Set(ID, Entity) error
}

