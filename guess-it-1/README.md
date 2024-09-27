guess-it-1
Objectives
The goal of this project is to build a program that, given a number as input, calculates and outputs a range in which the next number in the sequence is expected to be. This program leverages statistical and probability calculations to make its predictions, based on the input sequence.

Instructions
The program reads a series of numbers from standard input, which represent values on the y-axis of a graph, where the x-axis represents the line numbers (0, 1, 2, 3, 4, 5, ...). The objective is to guess the range for the next number in the sequence.

Example Input:
python
Copier le code
189
113
121
114
145
110
...
Example Output:
python
Copier le code
$ ./your_program
189 --> the standard input
120 200    --> the range for the next input (113)
113 --> the standard input
160 230    --> the range for the next input (121)
121 --> the standard input
110 140    --> the range for the next input (114)
114 --> the standard input
100 200    --> the range for the next input (145)
145 --> the standard input
1 99      --> the range for the next input (110)
Each line represents the predicted range for the next input value, separated by a space. The range can be adjusted or fine-tuned based on your calculations.

Implementation Details
The program must guess the range for the next number based on the previous inputs using statistical methods such as:

Linear regression
Probability distribution
Moving averages, etc.
Your solution should prioritize both accuracy and performance.

Testing
The program will be extensively tested with multiple datasets (Data 1, Data 2, and Data 3). The test will measure the accuracy of the predicted range for the next number in the sequence. The scoring system rewards smaller and more accurate ranges.

Steps for Testing:
Create a student/ folder.
Copy all necessary files to run your program into the student/ folder.
Create a shell script named script.sh that runs your program.
Example script.sh for a JavaScript solution:

sh
Copier le code
#!/bin/sh
# Navigate to the student folder and run the solution
node ./student/solution.js
You can write your program in one of the following languages:

Golang
JavaScript (JS)
Rust
Python
Make sure to follow these steps to ensure the testing process works correctly.

How to Run
Compile or run the program as appropriate for your chosen language (e.g., go run, python, node, etc.).
Input the numbers into the program, and it will output the predicted range for the next number in the sequence.
Project Benefits
By working on this project, you will:

Improve your skills in statistical and probability calculations.
Gain experience with scripting and automation.
