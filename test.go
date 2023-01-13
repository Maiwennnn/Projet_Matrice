package main

import (
	"fmt"
	"log"
	"os"
)

/*func check(e error) {
	if e != nil {
		panic(e)
	}
}*/

/*func ReadFile(fileName string) {

	fmt.Printf("\n\nReading a file in Go lang\n")
	//fileName := "test.txt"

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data) // conversion de byte en string
}

func WriteFile(text string, file *os.File) {
	file,err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
}*/

func main() {
	/*fmt.Println("J'aime Go")
	fmt.Println("Il est ", time.Now())
	ReadFile("test.txt")

	check(err)
	WriteFile("Brocolis\n", file)*/

	//x := 16
	//a := strconv.Itoa(x)

	/*err := os.WriteFile("test.txt", []byte(a), 0666)
	if err != nil {
		log.Fatal(err)
	}*/

	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	mat, err := os.Stdout.Write(data) // lis matrice dans fichier et l'écris dans terminal
	fmt.Println(" pause2")
	fmt.Println(mat)
}

// lecture et ecriture dans matrice en C
/*int lireMatrice(FILE *fich,int mat[N][N])
{int i,j,res;
 int val;
 for (i=0; i<N; i++)
   for (j=0; j<N; j++)
     {
       res=fscanf(fich, "%d",&val);
       if (res==EOF)
         {
           fprintf(stderr,"Fin de ficher atteinte: manque des coefficients\n");
           exit(-1);
         }
       mat[i][j]=val;
     }
 return(0);
}

int ecrireMatrice(FILE *fich,int mat[N][N])
{int i,j;
 for (i=0; i<N; i++)
   {
     for (j=0; j<N; j++)
       {
         fprintf(fich, "%d ",mat[i][j]);
       }
     fprintf(fich, "\n");
   }
 return(0);

 //

 ouverture fichier en C

 int  main(int argc,char *argv[])
{
  FILE *fichA, *fichB;
  char nomFichA[15]="matriceA.txt";
  char nomFichB[15]="matriceB.txt";

  int A[N][N], B[N][N], C[N][N];
  int err;

  fichA=fopen(nomFichA,"r");
  if (!fichA)
    {
      fprintf(stderr,"erreur d'ouverture du fichier %s\n",nomFichA);
      exit(-1);
    }

  fichB=fopen(nomFichB,"r");
  if (!fichB)
    {
      fprintf(stderr,"erreur d'ouverture du fichier %s\n",nomFichB);
      exit(-1);
    }

  err=lireMatrice(fichA,A);
  if (err)
    {
      fprintf(stderr,"erreur lors de la lecture de %s\n",nomFichA);
      exit(-1);
    }

  err=lireMatrice(fichB,B);
  if (err)
    {
      fprintf(stderr,"erreur lors de la lecture de %s\n",nomFichB);
      exit(-1);
    }

 fprintf(stdout," Matrice lues: \n");
 fprintf(stdout," A: \n");
 ecrireMatrice(stdout,A);

 fprintf(stdout," B: \n");
 ecrireMatrice(stdout,B);

 matProd(A,B,C);

 fprintf(stdout," resultat du produit C=AB (FAUX évidement): \n");
 ecrireMatrice(stdout,C);


   return(0);
}
}*/
