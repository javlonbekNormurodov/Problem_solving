import math
class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """
        sum = 0
        i = 5
        while i <= n:
            sum += n // i
            i *= 5
        return sum

print(Solution().trailingZeroes(6503))