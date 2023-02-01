from cmath import inf
from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxsum, current = -inf, 0

        for x in nums:
            current = max(x, current + x)
            maxsum = max(maxsum, current)

        return maxsum
