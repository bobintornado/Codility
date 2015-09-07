#Reasoning

This is the implementation of a classical Kadane's algorithm. 

For anyone read this algorithm, it's obvious that the key thing there is the the determine the max_ending at every element and when to make the `jump`. 

In the original algorithm, the jump point is 0, so the algorithm basically ignore any negative value, but this is variant, the max_ending is calculated by comparing the `element` itself and the `max_ending+element`. 

So in this way, the `jump` threshold is the element itself. 

This clearly cover the case where the sub-array has some negative values in the middle and a few very large positive numbers in both ends. 