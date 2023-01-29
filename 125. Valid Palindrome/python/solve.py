class Solution:
    def isPalindrome(self, s: str) -> bool:
        new = ''.join(filter(str.isalnum, s)).lower()

        if new == '':
            return True

        for i in range(0, len(new) // 2):
            if new[i] != new[len(new) - i - 1]:
                return False

        return True
