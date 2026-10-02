package main

import (
	"bytes"
	"time"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math"
	"math/rand"
	
	"robot/core"
	"github.com/google/uuid"
)

const (
	spacing = 50 // mm
	speed = 700 // mm/s
)

type Vector struct {
	X, Y float64
}

var x float64 = 0
var y float64 = 0
var angle float64 = 0

var epsilon float64 = 0.5

func GetRobotState(modeCh chan<- core.Mode, uuidCh chan<- uuid.UUID) {
	var response struct {
		RobotMode core.Mode	`json:"robot_mode"`
		MapId	uuid.UUID	`json:"map_id"`
	}
	
	var previousMode core.Mode
	for {
		resp, err := http.Get("http://localhost:8080/api/robot/mode")
		if err != nil {
			log.Fatal(err)
			time.Sleep(time.Second)
			continue
		}
		
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Println(err)
			resp.Body.Close()
			time.Sleep(time.Second)
			continue
		}


		resp.Body.Close()
		if response.RobotMode != previousMode {
			previousMode = response.RobotMode

			modeCh <- response.RobotMode
			uuidCh <- GetCurrentSession()
		}

		time.Sleep(time.Second)
	}
}

func GetCurrentSession() uuid.UUID {
	var response struct {
		MapId	uuid.UUID	`json:"map_id"`
	}
	
	resp, err := http.Get("http://localhost:8080/api/robot/session")
	if err != nil {
		log.Println(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	

	return response.MapId
}


func GeneratePoint() (float64, float64) {
	var rDirection int8

	turn := rand.Float64()
	turningRate := rand.Float64()

	if(turn < 0.05 && turn > 0.95) {
		rDirection = 0
	} else if turn > epsilon {
		epsilon += 0.01
		rDirection = -1
	} else {
		epsilon -= 0.01
		rDirection = 1
	}
	
	angle += float64(rDirection) * turningRate *  math.Pi/5;
	x += math.Cos(angle) * spacing
	y += math.Sin(angle) * spacing
	
	return x, y
}

func StopAutoMode(robot *core.Robot) {
	robot.SetMapId(uuid.Nil)
	x, y = 0.0, 0.0
}

func StartAutoMode(robot *core.Robot) {
	var vectors []Vector
	interval := time.Duration(spacing) * time.Second / time.Duration(speed)
	
	now := time.Now()

	vector := Vector{X: x, Y: y}
	vectors = append(vectors, vector)
	for {
		if robot.GetMode() != core.AUTO {
			StopAutoMode(robot)
			return
		}

		x, y = GeneratePoint()
		vector := Vector{X: x, Y: y}
		vectors = append(vectors, vector)
		
		if time.Since(now) >= 100 * time.Millisecond {
			var request struct {
				Vectors []Vector	`json:"vectors"`
				MapId uuid.UUID		`json:"map_id"`
			}

			request.Vectors = vectors
			request.MapId = robot.GetMapId()
			
			fmt.Println(request)
			body, err := json.Marshal(&request)
			if err != nil {
				log.Fatalf("Failed to marshal JSON: %s", err)
				continue
			}

			response, err := http.Post("http://localhost:8080/api/coordinates", "application/json", bytes.NewReader(body))
			if err != nil {
				log.Fatalf("POST failed: %s", err)
			}
			response.Body.Close()


			vectors = vectors[:0]
		}

		time.Sleep(interval)
	}
}

func main() {
	modeCh := make(chan core.Mode)
	uuidCh := make(chan uuid.UUID)
	robot := core.NewRobot()


	go GetRobotState(modeCh, uuidCh)

	for {
		mode := <-modeCh
		mapId := <-uuidCh
		
		robot.SetMode(mode)
		robot.SetMapId(mapId)

		if robot.GetMode() == core.AUTO {
			go StartAutoMode(robot)
		}

		fmt.Println("Mode changed:", robot.GetMode())
	}

}
