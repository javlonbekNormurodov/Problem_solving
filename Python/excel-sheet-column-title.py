class Solution(object):
    def convertToTitle(self, columnNumber):
        """
        :type columnNumber: int
        :rtype: str
        """
        NumberToSheet = {
            1: 'A',
            2: 'B',
            3: 'C',
            4: 'D',
            5: 'E',
            6: 'F',
            7: 'G',
            8: 'H',
            9: 'I',
            10: 'J',
            11: 'K',
            12: 'L',
            13: 'M',
            14: 'N',
            15: 'O',
            16: 'P',
            17: 'Q',
            18: 'R',
            19: 'S',
            20: 'T',
            21: 'U',
            22: 'V',
            23: 'W',
            24: 'X',
            25: 'Y',
            26: 'Z'
        }
        print(columnNumber // 26)
        if columnNumber // 26 > 0 :
            val = NumberToSheet[columnNumber // 26]
            while columnNumber > 26:
                columnNumber = columnNumber - 26
            val2 = columnNumber
            print(NumberToSheet[val] + NumberToSheet[val2])
            return val + NumberToSheet[val2]
        return NumberToSheet[columnNumber]

columnNumber = 28
print(Solution().convertToTitle(columnNumber))