#include "simil_check.h"

/*
 *  by Rei-chi Lin
 */

char* C_lib_version() 
{
	return "1.2.0";
}

double check_similarity(byte *b1, int *index, byte *b2, int len, int diff)
{
	*index = -1;
	double max_ratio = -1.0;
	for (int start_index = 0; start_index <= diff; start_index += 1) 
	{
		int count = 0;
		for(int i = 0; i < len; i += 1)
		{
			if(*(b1 + start_index + i) == *(b2 + i))
			{
				count += 1;
			}
		}
		double ratio = count * 100.0 / len;
		if ( ratio > max_ratio )
		{
			max_ratio = ratio;
			*index = start_index;
		}
	}
    return max_ratio;
}