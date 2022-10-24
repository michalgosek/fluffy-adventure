# What is a Queue?
A queue is linear data structure which models real world queue bys 
having two primary operations, namely **enqueue** and **dequeue**. 

There does not seem to be consistent terminology for inserting and removing 
elements from queues. 

Enqueue = Adding = Offering\
Dequeue = Polling 
 
# When and where is a queue used?

1) Any waiting line models a queue, for example a lineup at movie theatre. 
2) Can be used to efficiently keep track of the x most recently added elements. 
3) Web server request management where you want first come first serve. 
4) Breadth first search (BFS) graph traversal.

**Implementation**:

 

**Complexity Analysis**:

| List Type | Singly |
|-----------|--------|
| Enqueue   | O(1)   |
| Dequeue   | O(1)   |
| Peeking   | O(1)   | 
| Contains  | O(n)   |
| Removal   | O(n)   |
| Is Empty  | O(1)   |
