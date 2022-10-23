# What is a linked list?
A linked list is a sequential list of nodes that 
hold data which point to other nodes also containing data.
The last node has a null reference to next node.

# When and where is a linked list used?

1) Used in many List, Queue & Stack implementations. 
2) Great for creating circular lists. 
3) Can easily model real world objects such as trains. 
4) Used in separate chaining, which is present certain Hashtable 
implementations to deal with hashing collisions.
5) Often used in the implementation of adjacency lists for graphs. 

### Terminology
**Head**: The first node in a linked list.\
**Tail**: The last node in a linked list.\
**Pointer**: Reference to next node.\
**Node**: An object containing data and pointer(s)

# Singly vs Doubles Linked Lists

Singly linked lists only hold a reference to the next node. 
In the implementation you always maintain a reference to the **head** to the 
linked list and reference to the **tail** node for quick additions/removals. 

Doubly linked list is opposite to single linked list due to bidirectional character. 
Each node points to each other in two directions - next and previous. In the implementation you always maintain a reference to the **head** to the
linked list and reference to the **tail** node for quick additions/removals.


# Singly vs Doubles Linked Lists Pros and Cons 

| Type               | Pros                                    | Cons                                   |
|--------------------|-----------------------------------------|----------------------------------------|
| Single Linked List | Use less memory, simpler implementation | Cannot easily access previous elements |
| Doubly Linked List | Can be traversed backwards              | Takes 2x memory                        |


**Implementation**:

1. Create a static array with an initial capacity.
2. Add elements to the underlying static array, keeping track of the number of elements.
3. If adding another element will exceed the capacity, then create a new static array with
twice the capacity and copy the original elements into it. 


**Complexity Analysis**:

| List Type         | Singly | Doubly |
|-------------------|--------|--------|
| Search            | O(n)   | O(n)   |
| Insertion at head | O(1)   | O(1)   |
| Insertion at tail | O(1)   | O(1)   |
| Remove at head    | O(1)   | O(1)   |
| Remove at tail    | O(n)   | O(n)   |
| Remove in middle  | O(n)   | O(n)   |
