package shared

func ClampValue(value float32, minValue float32, maxValue float32) float32 {
    if value < minValue {
        return minValue
    } else if value > maxValue {
        return maxValue
    }

    return value
}
