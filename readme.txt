Projet multiplication de matrices :

Ce programme permet de multiplier deux matrices carrées de même taille qui sont dans les fichiers matriceA.txt et matriceB.txt dans le dossier client.
Le résultat de la multiplication est dans le fichier matriceC.txt qui apparaitra dans le dossier client.

Pour écrire correctement les matrices utilisées dans les fichiers matriceA.txt et matriceB.txt, il faut laisser un espace entre chaque nombre, sauf le dernier de la ligne.
Il faut simplement revenir a la ligne après avoir écrit le dernier nombre de la ligne.

Pour run le programme, il faut run en premier dans un terminal server_tcp.go et ensuite dans un autre terminal client_tcp.go

Lorsque le programme affiche "File received" dans le terminal du client, la matrice a été calculé et le programme est fini.

Nous aurions aimé trouver un moyen d'envoyer les deux fichiers avec une seule connection TCP mais nous n'avons pas réussi.
Par conséquent, le client se déconnecte après l'envoi du premier fichier puis se reconnecte pour envoyer le deuxième fichier. 
Il se reconnecte ensuite une troisième fois pour attendre la réponse. 
