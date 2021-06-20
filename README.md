# Go Data Collector Test

Testing various Go packages and their support for collecting system information

## Purpose

The purpose of this project is to determine what the best library there is to pull system information on a variety of systems. After completing some tests, I will be picking a library which I will be using in an upcoming health monitoring project.

## Criteria

I will only be testing on Mac but the ideal library will have support for Windows/Linux/Mac (any additional OS support is a bonus). Linux support is the most crucial.

Each library should have extensive support for getting basic system information (CPU, OS, Total Memory, etc.) as well as extensive support for getting more detailed system information (processes, network info., memory usage, etc.)

## Libraries being tested

[gopsutil](https://github.com/shirou/gopsutil)

[go-sysinfo](https://github.com/elastic/go-sysinfo)

## Notes

- This code is sloppy but it serves a purpose. The extra time to cleanup the code simply wasn't worth it for a project like this.
- The data isn't one-to-one mapped so it requires a bit more investigation when trying to figure out what's supported via each library.
