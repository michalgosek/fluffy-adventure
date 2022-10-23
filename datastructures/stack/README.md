# What is a stack?

A stack is a one-endend linear data structure which models a real 
world stuck by having two primary operations, namely **push** and **pop**. 

# When and where is a stack used?

1) Used by undo mechanism in text editors.
2) Used in compiler syntax checking for matching brackets and braces.
3) Can be used to model a pile of books or plates.
4) Used behind the scenes to support recursion by keeping track of previous function calls.
5) Can be used to do a Depth First Search (DFS) on a graph. 

 
**Implementation**:

1) Create a static array with an initial capacity.
2) Add elements to the underlying static array, keeping track of the number of elements.
3) If adding another element will exceed the capacity, then create a new static array with
twice the capacity and copy the original elements into it. 

**Complexity Analysis**: 

| Operation       | Time |
|-----------------|------|
| Pushing         | O(1) |
| Popping at head | O(1) |
| Peeking         | O(1) |
| Searching       | O(n) |
| Size            | O(1) |
