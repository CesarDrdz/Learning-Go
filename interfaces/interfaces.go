//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Sedan string
type Truck string

// * Vehicles have a model name in addition to the vehicle type:
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

// * Vehicles have a model name in addition to the vehicle type:
func (s Sedan) String() string {
	return fmt.Sprintf("Sedan: %v", string(s))
}

// * Vehicles have a model name in addition to the vehicle type:
func (t Truck) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(t))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (s Sedan) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

//   - Write a single function to handle all of the vehicles
//     that the shop works on.
func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to the Small Lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to the Standard Lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to the Large Lift\n", p)
	}
}

//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name

func main() {
	sedan := Sedan("Sporty")
	truck := Truck("Work")
	motocycle := Motorcycle("Chopper")
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	sendToLift(sedan)
	sendToLift(truck)
	sendToLift(motocycle)
}
