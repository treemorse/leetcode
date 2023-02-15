from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        prev = ListNode(0)
        prev.next = head
        record = prev
        trav = head
        while trav:
            if trav.val == val:
                prev.next = trav.next
                trav = prev.next
            else:
                prev = trav
                trav = trav.next
        return record.next