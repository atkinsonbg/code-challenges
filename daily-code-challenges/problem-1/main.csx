#!/usr/bin/env dotnet-script

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

var numbers = Args[0].Split(',').Select(int.Parse).ToList();
var k = int.Parse(Args[1]);

var exists = false;
var sumNumbers = new int[2];

var seenNumbers = new List<int>();
foreach (var number in numbers)
{
    int complement = k - number;
    if (seenNumbers.Contains(complement))
    {
        sumNumbers[0] = complement; 
        sumNumbers[1] = number;
        exists = true;
        break;
    }
    seenNumbers.Add(number);
}
 
Console.WriteLine(exists);
Console.WriteLine(string.Join(", ", sumNumbers));
