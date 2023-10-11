# What does the below code snippet print out? How can we fix the anonymous functions to behave as we'd expect?

functions = []
for i in range(10):
    functions.append(lambda i=i: i)

for f in functions:
    print(f())

# this is due to how closures capture variables in Python. In the first loop, you're creating a list of
# lambda functions, and each of these functions returns the value of i. However, these lambda functions do
# not capture the current value of i when they are defined. Instead, they store a reference to the variable
# i, and will look up its value when they are called. By the time you call the functions in the second loop,
# the first loop has already finished executing, and i has been set to its final value of 9. Therefore, when
# any of the lambda functions are called, they all return 9, which is the final value of i. In this modified code,
# i=i in the lambda function definition sets the default value of i for each lambda function to be the current
# value of i in the loop. This way, each function will return the expected value from 0 to 9.

# Links
# https://discuss.python.org/t/make-lambdas-proper-closures/10553
#