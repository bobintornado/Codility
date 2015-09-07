#Reasoning

This is the implementation of a classical Kadane's algorithm with some tweaks, which handle negative values and single element input (assuming the algorithm couldn't return empty array).

For anyone read this algorithm, it's obvious that the key thing step is the the step to determine the max_ending at every element and when to make the `jump`. 

In the original algorithm, the jump point is `0`, so the algorithm basically ignore any negative value, but in this variant, the max_ending is calculated by comparing the `element` itself and the `max_ending+element`. 

So in this way, the `jump` threshold is the element itself. 

This clearly cover the case where the sub-array has some negative values in the middle and a few very large positive numbers in both ends. 

Reference : <https://en.wikipedia.org/wiki/Maximum_subarray_problem>