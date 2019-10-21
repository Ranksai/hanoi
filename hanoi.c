#include <stdio.h>
#include <stdlib.h>

int count = 0;

typedef struct hanoi
{
    int[] 
}hanoi;


void replace(int plate, int *from, int *to){
    int loop;
    int i;
    loop = sizeof from / sizeof from[0];
    count++;
    printf("%d\n", count);
    for (i = 0; i < loop; i++)
    {    
        printf("Step %d %d, moves %s to %c\n", count, plate, from, to);
    
    }
}

void hanoi02(int plate, int *from, int *to, int *rest){
    if (plate > 1) {
        hanoi02(plate-1, from, rest, to);
        // replace(plate, from, to);
        printf("Step %d %d, moves %s to %c\n", count, plate, from, to);
        hanoi02(plate-1, to, rest, from);
    } else {
        replace(plate, from, to);
    }
}

void hanoi01(int plate, char *from, char *to, char *rest)
{
    if (plate > 1)
    {
        hanoi01(plate - 1, from, rest, to);
        // replace(plate, from, to);
        count++;
        printf("Step %d %d moves %s to %s\n", count, plate, from, to);
        hanoi01(plate - 1, to, rest, from);
    }
    else
    {
        count++;
        printf("Step %d %d moves %s to %s\n", count, plate, from, to);
    }
}

int main (int argc, char *argv[]){

    int plate;
    int i;
    char *fromC;
    char *toC;
    char *restC;
    char *endl;
    int *fromI;
    int *toI;
    int *restI;

    plate = strtol(argv[1], &endl, 10);
    fromC = "A";
    toC =  "B";
    restC = "C";
    fromI = (int*)malloc(sizeof(int)*plate);
    toI = (int*)malloc(sizeof(int)*plate);
    restI = (int*)malloc(sizeof(int)*plate); 

    for(i = 0; i < plate; i++) {
        fromI[i] =  i+1;
    }

    hanoi01(plate, fromC, toC, restC);

    free(fromI);
    free(toI);
    free(restI);
}
