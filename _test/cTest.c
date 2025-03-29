#include <stdio.h>
#include <stdlib.h>
#include <time.h>

/* Constants for the program */
#define MAX_SIZE 10
#define MIN_VALUE 1
#define MAX_VALUE 100

/**
 * Function to generate an array of random integers
 * 
 * @param arr The array to fill
 * @param size Size of the array
 */
void generateRandomArray(int arr[], int size) {
    // Seed the random number generator
    srand(time(NULL));
    
    // Fill the array with random values
    for (int i = 0; i < size; i++) {
        arr[i] = MIN_VALUE + rand() % (MAX_VALUE - MIN_VALUE + 1);
    }
}

/**
 * Function to find the maximum value in an array
 * 
 * @param arr The array to search
 * @param size Size of the array
 * @return The maximum value found
 */
int findMax(int arr[], int size) {
    int max = arr[0];  // Start with first element as max
    
    /* Loop through the array to find the maximum value */
    for (int i = 1; i < size; i++) {
        if (arr[i] > max) {
            max = arr[i];  // Update max if current element is larger
        }
    }
    
    return max;
}

/**
 * Function to find the minimum value in an array
 * 
 * @param arr The array to search
 * @param size Size of the array
 * @return The minimum value found
 */
int findMin(int arr[], int size) {
    int min = arr[0];  // Start with first element as min
    
    // Loop through the array to find the minimum value
    for (int i = 1; i < size; i++) {
        if (arr[i] < min) {
            min = arr[i];  // Update min if current element is smaller
        }
    }
    
    return min;
}

/**
 * Function to calculate the average of array elements
 * 
 * @param arr The array of integers
 * @param size Size of the array
 * @return The average value
 */
float calculateAverage(int arr[], int size) {
    int sum = 0;
    
    // Calculate sum of all elements
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    
    // Return average (cast to float for decimal result)
    return (float)sum / size;
}

/**
 * Main function
 */
int main() {
    int data[MAX_SIZE];
    int size = MAX_SIZE;
    
    // Generate random data
    generateRandomArray(data, size);
    
    // Print the array
    printf("Generated array: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", data[i]);
    }
    printf("\n");
    
    // Find and print statistics
    int max = findMax(data, size);
    int min = findMin(data, size);
    float avg = calculateAverage(data, size);
    
    printf("Statistics:\n");
    printf("- Maximum value: %d\n", max);
    printf("- Minimum value: %d\n", min);
    printf("- Average value: %.2f\n", avg);
    
    return 0;
}