package main

import (
    "errors"
    "regexp"
    "time"
)


func validateStringField(field string, pattern string) bool {
    matched, _ := regexp.MatchString(pattern, field)
    return matched
}

func validateReceipt(receipt *Receipt) error {
    // Validate retailer name 
    if !validateStringField(receipt.Retailer, `^[\w\s\-&]+$`) {
        return errors.New("invalid retailer format")
    }

    // Validate purchaseDate (format: YYYY-MM-DD)
    layout := "2006-01-02"
    _, err := time.Parse(layout, receipt.PurchaseDate)
    if err != nil {
        return errors.New("invalid purchaseDate format")
    }

    // Validate purchaseTime (format: HH:MM in 24-hour format)
    layoutTime := "15:04" 
    _, errTime := time.Parse(layoutTime, receipt.PurchaseTime)
    if errTime != nil {
        return errors.New("invalid purchaseTime format")
    }

    // Validate total (monetary value with two decimal places)
    totalPattern := `^\d+\.\d{2}$`
    if !validateStringField(receipt.Total, totalPattern) {
        return errors.New("invalid total format")
    }

    // Validate each item
    for _, item := range receipt.Items {
        // ShortDescription: allows word characters, whitespace, hyphens, and ampersands
        if !validateStringField(item.ShortDescription, `^[\w\s\-&]+$`) {
            return errors.New("invalid item shortDescription format")
        }
        // Price: monetary value with two decimal places
        if !validateStringField(item.Price, `^\d+\.\d{2}$`) {
            return errors.New("invalid item price format")
        }
    }

    return nil
}

