import re

class Solution:
    def removeVowels(self, S: str) -> str:
        regexp_vowels = "(a|i|u|e|o)"
        return re.sub(regexp_vowels, "", S)
