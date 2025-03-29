package com.example.demo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

/**
 * A simple calculator class to demonstrate Java comments
 * @author Example Developer
 */
public class Calculator {
    
    // Constants for special calculations
    private static final double PI = 3.14159;
    private static final double E = 2.71828;
    
    // List to store calculation history
    private List<String> history;
    
    /**
     * Constructor for Calculator
     */
    public Calculator() {
        this.history = new ArrayList<>();
    }
    
    /**
     * Add two numbers
     * @param a first number
     * @param b second number
     * @return sum of the numbers
     */
    public double add(double a, double b) {
        double result = a + b;
        history.add(a + " + " + b + " = " + result);
        return result;
    }
    
    /**
     * Subtract second number from first
     * @param a first number
     * @param b second number
     * @return difference of the numbers
     */
    public double subtract(double a, double b) {
        double result = a - b;
        history.add(a + " - " + b + " = " + result);
        return result;
    }
    
    /**
     * Multiply two numbers
     * @param a first number
     * @param b second number
     * @return product of the numbers
     */
    public double multiply(double a, double b) {
        // Perform multiplication
        double result = a * b;
        history.add(a + " * " + b + " = " + result);
        return result;
    }
    
    /**
     * Divide first number by second
     * @param a first number (dividend)
     * @param b second number (divisor)
     * @return quotient of the division
     * @throws IllegalArgumentException if divisor is zero
     */
    public double divide(double a, double b) {
        // Check for division by zero
        if (b == 0) {
            throw new IllegalArgumentException("Cannot divide by zero");
        }
        
        double result = a / b;
        history.add(a + " / " + b + " = " + result);
        return result;
    }
    
    /**
     * Calculate area of a circle
     * @param radius radius of the circle
     * @return area of the circle
     */
    public double areaOfCircle(double radius) {
        if (radius < 0) {
            throw new IllegalArgumentException("Radius cannot be negative");
        }
        
        /* 
         * The formula for the area of a circle is:
         * A = π * r²
         */
        double result = PI * radius * radius;
        history.add("Area of circle with radius " + radius + " = " + result);
        return result;
    }
    
    /**
     * Get calculation history
     * @return list of calculation strings
     */
    public List<String> getHistory() {
        return new ArrayList<>(history);
    }
    
    /**
     * Clear calculation history
     */
    public void clearHistory() {
        history.clear();
    }
    
    /**
     * Search history for specific term
     * @param term search term
     * @return filtered history list
     */
    public List<String> searchHistory(String term) {
        // Filter history items containing the search term
        return history.stream()
            .filter(entry -> entry.contains(term))
            .collect(Collectors.toList());
    }
    
    // Main method for testing
    public static void main(String[] args) {
        Calculator calc = new Calculator();
        
        // Perform some calculations
        System.out.println("5 + 3 = " + calc.add(5, 3));
        System.out.println("10 - 4 = " + calc.subtract(10, 4));
        System.out.println("6 * 7 = " + calc.multiply(6, 7));
        System.out.println("20 / 5 = " + calc.divide(20, 5));
        System.out.println("Area of circle with radius 3 = " + calc.areaOfCircle(3));
        
        // Print history
        System.out.println("\nCalculation History:");
        calc.getHistory().forEach(System.out::println);
    }
}