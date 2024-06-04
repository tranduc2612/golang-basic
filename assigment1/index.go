package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CREATE string = "create"
	STATUS string = "status"
	PARK   string = "park"
	LEAVE  string = "leave"
	HELP   string = "help"
	EXIT   string = "exit"
)

var current_slot []ParkSlot

func renderCommand() {
	fmt.Println("-----------------------------List command you can use--------------------------------")
	fmt.Println("status : status of parking")
	fmt.Println("create <num>  : to create a new slot for parking")
	fmt.Println("park <arg> : to create tickets for parking")
	fmt.Println("leave <arg> <hours> : to leave tickets for parking")
	fmt.Println("exit")
}

func handleParkingCar(numberRegsiter string) {
	for i := 0; i < len(current_slot); i++ {
		if current_slot[i].current_ticket == nil {
			newCar := NewCar(numberRegsiter)
			newTicket := NewTicket(newCar, 0, &current_slot[i])
			current_slot[i].AddTicket(newTicket)
			fmt.Println("--------------------------------Parking successfully------------------------")
			return
		}
	}
	fmt.Println("--------------------------------Parking is full slot------------------------")
}

func handleLeave(numberRegsiter string, hours string) {
	convertHours, err := strconv.ParseFloat(hours, 64)
	for i := 0; i < len(current_slot); i++ {
		if current_slot[i].current_ticket.car.registation_number == numberRegsiter && err == nil {
			// handle leaving the car
			totalPrice := current_slot[i].current_ticket.Pay(convertHours)
			slotId := strconv.Itoa(current_slot[i].id)
			current_slot[i].current_ticket.Pay(convertHours)
			current_slot[i].SumTotal()
			fmt.Println("Regesitration Number " + numberRegsiter + " from Slot " + slotId + " has left with time " + hours + " and total money " + strconv.FormatFloat(totalPrice, 'f', 2, 64))
			current_slot[i].current_ticket = nil
			return
		}
	}
	if err != nil {
		fmt.Println("Invalid hours number")
	} else {
		fmt.Println("\nRegesitration Number" + numberRegsiter + "not found")

	}
	fmt.Println("--------------------------------Ticket is invalid------------------------")
}

func handleCreateSlot(input string) {
	number, err := strconv.Atoi(input)
	if err != nil || number < 0 {
		fmt.Println("number is invalid")
	} else {
		current_slot = make([]ParkSlot, number)

		for i := 0; i < number; i++ {
			current_slot[i] = *NewParking(i+1, 0)
		}
		fmt.Println("--------------------------------Create Slot successfully------------------------")
	}
}

// Hàm để tìm độ dài lớn nhất của mỗi cột
func findMaxColumnWidths(table [][]string) []int {
	if len(table) == 0 {
		return []int{}
	}
	columnWidths := make([]int, len(table[0]))
	for _, row := range table {
		for colIdx, col := range row {
			if len(col) > columnWidths[colIdx] {
				columnWidths[colIdx] = len(col)
			}
		}
	}
	return columnWidths
}

// Hàm để tạo ra một dòng kẻ
func createSeparator(columnWidths []int) string {
	var separator strings.Builder
	separator.WriteString("+")
	for _, width := range columnWidths {
		separator.WriteString(strings.Repeat("-", width+2) + "+")
	}
	return separator.String()
}

// Hàm để in bảng
func printTable(table [][]string) {
	if len(table) == 0 {
		return
	}

	columnWidths := findMaxColumnWidths(table)
	separator := createSeparator(columnWidths)

	fmt.Println(separator)
	for _, row := range table {
		fmt.Print("|")
		for colIdx, col := range row {
			fmt.Printf(" %-*s |", columnWidths[colIdx], col)
		}
		fmt.Println()
		fmt.Println(separator)
	}
}

func run() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Kiểm tra độ dài của mảng split trước khi truy cập các phần tử
		split := strings.Fields(input)
		switch split[0] {
		case CREATE:
			if len(split) > 1 {
				handleCreateSlot(split[1])
			} else {
				fmt.Println("create command requires an additional argument")
			}
		case PARK:
			if len(split) > 1 {
				handleParkingCar(split[1])
			} else {
				fmt.Println("create command requires an additional argument")
			}
		case LEAVE:
			if len(split) > 2 {
				handleLeave(split[1], split[2])
			} else {
				fmt.Println("create command requires an additional argument")
			}
		case STATUS:
			table := [][]string{
				{"ID_SLOT", "MONEY", "STATUS", "REGISTRATION NUMBER"},
			}
			for i := 0; i < len(current_slot); i++ {
				var numberRegsiter string = ""
				if current_slot[i].current_ticket != nil {
					numberRegsiter = current_slot[i].current_ticket.car.registation_number
				}
				table = append(table, []string{current_slot[i].GetId(), current_slot[i].GetMoney(), current_slot[i].GetStatus(), numberRegsiter})
			}
			printTable(table)
		case HELP:
			renderCommand()
		case EXIT:
			fmt.Println("-----------------------------Thank you for using my service ! <3--------------------------------")
			return
		default:
			fmt.Println("command invalid")
		}
	}
}

func main() {

	renderCommand()
	fmt.Print("\n")
	fmt.Print("-----------------------------Welcome to Parking--------------------------------\n")
	run()

}
