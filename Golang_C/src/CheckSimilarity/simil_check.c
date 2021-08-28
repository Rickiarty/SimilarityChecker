#include "simil_check.h"

/*
 *  by Rei-chi Lin
 */

char* C_lib_version() 
{
	return "1.3.0";
}

int check_similarity(byte *b1, int *index, byte *b2, int len, int diff)
{
	*index = -1;
	double max_count = -1;
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
		if ( count > max_count )
		{
			max_count = count;
			*index = start_index;
		}
	}
    return max_count;
}

char* C_lib_author() 
{
	return "Rei-chi Lin";
}
