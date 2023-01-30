class Solution:
    def search(self, nums: list[int], target: int) -> int:
        low = 0
        high = len(nums) - 1

        while low <= high:
            middle = (low + high) // 2
            guess = nums[middle]

            if guess == target:
                return middle

            if guess > target:
                high = middle - 1

            if guess < target:
                low = middle + 1

        return -1
