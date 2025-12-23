/*
========================================================================
   USAGE & COMPILATION INSTRUCTIONS (Linux & Windows)
   Prerequisite: Make sure Go (Golang) is installed on your system.
========================================================================

1. COMPILATION (Build the Binary):
   Run the following command in your terminal:
   Command: go build -o port_scanner main.go
   (This will create an executable binary file named 'port_scanner')

2. EXECUTION (Run the Tool):
   - Linux/Mac: ./port_scanner
   - Windows:   .\port_scanner.exe

3. IMPORTANT NOTES:
   - Target: The target is currently hardcoded to "scanme.nmap.org".
     To scan a different IP, modify the 'addr' variable inside the code.
   - Concurrency: This tool launches 10,000 concurrent goroutines.
     It is extremely fast but might overwhelm slower networks.

4. DISCLAIMER:
   This tool is developed for educational purposes only. 
   Do not use it on targets without proper authorization.
========================================================================
*/
