class Node:
    def __init__(self, val, next_node=None):
        self.val = val
        self.next_node = next_node

def print_singly_linked_list(head):
    tmp = head
    while tmp != None:
        print(tmp.val, end=' ')
        tmp = tmp.next_node

def add_to_singly_linked_list(head, val):
    tmp = head
    while tmp.next_node != None:
        tmp = tmp.next_node
    tmp.next_node = Node(val)
"""
For your reference, here's the node structure:
class Node:
    def __init__(self, val, next_node=None):
        self.val = val
        self.next_node = next_node
"""

def remove_all_n(head, n):
    # Fill in the logic here
    # Store head node
    temp = head
    prev = None

    # If head node itself holds the key
    # or multiple occurrences of key
    while (temp != None and temp.val == n):
        head = temp.next_node  # Changed head
        temp = head         # Change Temp

    # Delete occurrences other than head
    while (temp != None):

        # Search for the key to be deleted,
        # keep track of the previous node
        # as we need to change 'prev.next'
        while (temp != None and temp.val != n):
            prev = temp
            temp = temp.next_node

        # If key was not present in linked list
        if (temp == None):
            return head

        # Unlink the node from linked list
        prev.next_node = temp.next_node

        # Update Temp for next iteration of outer loop
        temp = prev.next_node
    return head
def main():
    inp = input().split()
    n = inp.pop(0)
    size = inp.pop(0)
    head = Node(inp.pop(0))
    while inp: add_to_singly_linked_list(head, inp.pop(0))
    head = remove_all_n(head, n)
    print_singly_linked_list(head)

if __name__=="__main__":
    main()