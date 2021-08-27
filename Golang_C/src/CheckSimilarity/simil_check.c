#include "simil_check.h"

/*
 *  by Rei-chi Lin
 */

char* C_lib_version() 
{
	return "1.0.0";
}

double check_similarity(byte *b1, int start_index, byte *b2, int len)
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
    return ratio;
}