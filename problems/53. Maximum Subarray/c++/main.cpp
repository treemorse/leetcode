class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int current = 0, maxsum = INT_MIN;

        for(auto x : nums)
            current = max(x, current + x), \
            maxsum = max(maxsum, current);

        return maxsum;
    }
};