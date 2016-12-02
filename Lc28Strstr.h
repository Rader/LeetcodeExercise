#pragma once

#include <string>
using namespace std;

namespace leetcode {
	class Solution
	{
	public:
		int strStr(string haystack, string needle) {
			auto strlen = haystack.length();
			auto substrlen = needle.length();
			if (strlen < substrlen)
				return -1;

			size_t i, j;
			for (i = 0; i < strlen - substrlen; i++)
			{
				for (j = 0; j <= substrlen; j++)
				{
					if (haystack[i + j] != needle[j])
						break;
				}

				if (j == substrlen)
					return i;
			}
			return -1;
		}
	};

}
