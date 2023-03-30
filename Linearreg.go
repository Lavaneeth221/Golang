// loadCarData loads the car dataset from a CSV file and returns it as a slice of data points.
func loadCarData(car_data string) ([]*DataPoint, error) {
	f, err := os.Open(car_data)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []*DataPoint

	for _, record := range records {
		if len(record) != 7 {
			continue
		}

		var dp DataPoint
		dp.Features = make([]float64, 6)

		for i := 0; i < 6; i++ {
			switch record[i] {
			case "low":
				dp.Features[i] = 0
			case "med":
				dp.Features[i] = 1
			case "high":
				dp.Features[i] = 2
			case "vhigh":
				dp.Features[i] = 3
			default:
				return nil, fmt.Errorf("invalid value: %s", record[i])
			}
		}

		switch record[6] {
		case "unacc":
			dp.Label = 0
		case "acc":
			dp.Label = 1
		case "good":
			dp.Label = 2
		case "vgood":
			dp.Label = 3
		default:
			return nil, fmt.Errorf("invalid value: %s", record[6])
		}

		data = append(data, &dp)
	}

	return data, nil
}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Load the car dataset
	carData, err := loadcarData("car_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Shuffle the data randomly
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(carData), func(i, j int) {
		carData[i], carData[j] = carData[j], carData[i]
	})

	// Split the data into training and testing sets
	trainData, testData := splitData(carData, 0.8)

	// Convert the data into matrices
	XTrain, yTrain := dataToMatrices(trainData)
	XTest, yTest := dataToMatrices(testData)

	// Train the model on the training data
	model := Train(XTrain, yTrain)

	// Evaluate the model on the testing data
	accuracy := Evaluate(model, XTest, yTest)

	fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
}

// loadcarData loads the car dataset from a CSV file and returns it as a slice of data points.
func loadcarData(filename string) ([]*DataPoint, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []*DataPoint

	for _, record := range records {
		if len(record) != 5 {
			continue
		}

		var dp DataPoint
		dp.Features = make([]float64, 4)

		for i := 0; i < 4; i++ {
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return nil, err
			}

			dp.Features[i] = val
		}

		if record[4] == "car-setosa" {
			dp.Label = 0
		} else {
			dp.Label = 1
		}

		data = append(data, &dp)
	}

	return data, nil
}

// splitData splits the data into training and testing sets.
func splitData(data []*DataPoint, trainRatio float64) ([]*DataPoint, []*DataPoint) {
	trainSize := int(float64(len(data)) * trainRatio)
	return data[:trainSize], data[trainSize:]
}

// dataToMatrices converts the data into matrices.
func dataToMatrices(data []*DataPoint) (*mat.Dense, *mat.Dense) {
	var XData []float64
	var yData []float64

	for _, dp := range data {
		XData = append(XData, dp.Features...)
		yData = append(yData, dp.Label)
	}

	X := mat.NewDense(len(data), 4, XData)
	y := mat.NewDense(len(data), 1, yData)

	return X, y
}

// DataPoint represents a data point in the car dataset.
type DataPoint struct {
	Features []float64
	Label    float64
}

// Train trains a linear regression model on the given training data.
func Train(X *mat.Dense, y *mat.Dense) *mat.Dense {
	nSamples, n
