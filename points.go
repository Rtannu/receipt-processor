package main

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type PointsResponse struct {
    Points int `json:"points"`
}

// One point for every alphanumeric character in the retailer name.
func pointsForRetailerName(name string) int {
    alphanumericRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
    return len(alphanumericRegex.FindAllString(name, -1))
}

// 50 points if the total is a round dollar amount with no cents.
func pointsForRoundTotal(total string) int {
    if strings.HasSuffix(total, ".00") {
        return 50
    }
    return 0
}

// 25 points if the total is a multiple of 0.25.
func pointsForMultipleOfQuarter(total string) int {
    totalFloat, err := strconv.ParseFloat(total, 64)
    if err == nil && math.Mod(totalFloat, 0.25) == 0 {
        return 25
    }
    return 0
}

// 5 points for every two items on the receipt.
func pointsForItemPairs(items []Item) int {
    return (len(items) / 2) * 5
}


// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func pointsForDescriptionLength(items []Item) int {
    points := 0
    for _, item := range items {
        trimmedDesc := strings.TrimSpace(item.ShortDescription)
        if len(trimmedDesc)%3 == 0 {
            priceFloat, err := strconv.ParseFloat(item.Price, 64)
            if err == nil {
                // Always round up to the nearest integer
                additionalPoints := math.Ceil(priceFloat * 0.2)
                points += int(additionalPoints)
            }
        }
    }
    return points
}

// 6 points if the day in the purchase date is odd.
func pointsForOddDay(purchaseDate string) int {
    date, err := time.Parse("2006-01-02", purchaseDate)
    if err == nil && date.Day()%2 != 0 {
        return 6
    }
    return 0
}

// 10 points if the purchase time is between 2:00pm and 4:00pm.
func pointsForPurchaseTime(purchaseTimeStr string) int {
    purchaseTime, err := time.Parse("15:04", purchaseTimeStr)
    if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
        return 10
    }
    return 0
}

