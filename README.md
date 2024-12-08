# Temperature Converter

A simple Go program that converts temperatures between Celsius (C), Fahrenheit (F), and Kelvin (K). 

## Features
- Converts temperatures between:
  - Celsius (C)
  - Fahrenheit (F)
  - Kelvin (K)
- Accepts input from the user interactively.

## How It Works
1. Enter the temperature value to convert.
2. Specify the unit of the input temperature (`C`, `F`, or `K`).
3. Specify the unit to convert the temperature to (`C`, `F`, or `K`).
4. The program calculates and displays the converted temperature.

## Example Usage
```plaintext
Enter the Temperature: 100
From Unit (C, F, K) : C
To Unit (C, F, K) : F
100 C is 212.00 F
```
![Screenshot 2024-12-08 193817](https://github.com/user-attachments/assets/4a342178-afe2-46e1-9dc8-176acae1924a)


## Installation and Execution
### Prerequisites
- Install [Go](https://golang.org/dl/) on your system.

### Steps to Run
1. Clone this repository:
   ```bash
   git clone https://github.com/Jeeban369/Temperature-Converter.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Temperature-Converter
   ```
3. Run the program:
   ```bash
   go run main.go
   ```

## Code Explanation
- The program takes user input for:
  - The temperature value.
  - The source unit (`C`, `F`, `K`).
  - The target unit (`C`, `F`, `K`).
- It uses a `switch` statement to determine the conversion logic based on the source and target units.
- Outputs the converted temperature with 2 decimal precision.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
