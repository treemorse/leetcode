# very bad solution o(n^2)
# class Solution:
#     def twoSum(self, nums: list[int], target: int) -> list[int]:
#         hashmap = {}
#
#         for i in range(0, len(nums) - 1):
#             for j in range(i + 1, len(nums)):
#                 hashmap[nums[i]+nums[j]] = [i, j]
#
#         return hashmap.get(target)

class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        hashmap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap:
                return [i, hashmap[complement]]
            hashmap[nums[i]] = i