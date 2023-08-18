impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let (mut l, mut r) = (0, nums.len());

        while l < r {
            let mid = l + (r - l) / 2;
            let guess: i32 = nums[mid];

            if guess == target {
                return mid as _
            }

            if guess > target {
                r = mid
            }

            if guess < target {
                l = mid + 1
            }
        }

        return -1
    }
}