# TaxCalculator
Interview Technical Assessment to calculate tax totals and brackets for Canadian Federal Tax based on a salary input

This application is my solution to the assessment here: https://github.com/points/interview-test-server
Which assumes you've started the server as instructed there.

## Running the API

1. Navigate to the root dir of the project
2. Run the following command `make run`
3. Access the api with endpoints `/tax-calculator` and `/brackets` replacing query parameter values
    - `http://localhost:8080/tax-calculator?year=2022&salary=17`
    - `http://localhost:8080/brackets?year=2022`

## Running tests

1. Navigate to the root dir of the project
2. Run the following command `make test`


# API Documentation

## Get Tax Brackets for tax year

Request Method: `GET`
Endpoint: `/brackets`
Query Parameters:
    - `year`: tax year for retrieving brackets. Format: YYYY
    
Request Sample:
`http://localhost:8080/brackets?year=2022`

Response Sample:
```
{
    tax_brackets: {
        tax_brackets: [
        {
            min: 0,
            max: 50197,
            rate: 0.15
        },
        {
            min: 50197,
            max: 100392,
            rate: 0.205
        },
        {
            min: 100392,
            max: 155625,
            rate: 0.26
        },
        {
            min: 155625,
            max: 221708,
            rate: 0.29
        },
        {
            min: 221708,
            max: 0,
            rate: 0.33
        }]
    }
}
```

## Calculate Taxes

Request Method: `GET`
Endpoint: `/tax-calculator`
Query Parameters:
    - `year`: tax year for retrieving brackets. Format: YYYY
    - `salary`: annual income to be used in tax calculation
    
Request Sample:
`http://localhost:8080/tax-calculator?year=2022&salary=17`

Response Sample:
```
{
    tax_calculation: {
        total: 2.55,
        effective_rate: 0.15,
        tax_brackets: [
        {
            min: 0,
            max: 50197,
            rate: 0.15,
            actual: 17,
            tax: 2.55
        },
        {
            min: 50197,
            max: 100392,
            rate: 0.205,
            actual: 0,
            tax: 0
        },
        {
            min: 100392,
            max: 155625,
            rate: 0.26,
            actual: 0,
            tax: 0
        },
        {
            min: 155625,
            max: 221708,
            rate: 0.29,
            actual: 0,
            tax: 0
        },
        {
            min: 221708,
            max: 0,
            rate: 0.33,
            actual: 0,
            tax: 0
        }]
    }
}
```
