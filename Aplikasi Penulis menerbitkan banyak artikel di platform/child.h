#ifndef CHILD_H_INCLUDED
#define CHILD_H_INCLUDED
#include <iostream>
#define next(P) P->next
#define prev(P) P->prev
#define first(L) L.first
#define last(L) L.last
#define info(P) P->info
#define child(P) P->child
#define parent(P) P->parent
#include "parent.h"
using namespace std;
struct List_relasi;
struct artikel{
    string judulArtikel;
    int tahunTerbit;
    string konten;
};
struct elmlist_child;


typedef artikel infotype_child;
typedef struct elmlist_child *address_child;
struct List_child{
    address_child first;
    address_child last;
};
struct elmlist_child{
    infotype_child info;
    address_child next;
    address_child prev;
    struct elmlist_parent *parent;
};

void createList(List_child &L);
void insertFirst(List_child &L, address_child P);
void deleteFirst(List_child &L, address_child &P);
address_child alokasiChild(infotype_child x);
void dealokasi(address_child &P);
address_child findElm_judul(List_child L, string x);
void searchElm_judul(List_child L, string x);
void searchElm_tahun(List_child L, int x);
void searchElm_konten(List_child L, string x);
void printInfo(List_child L);
void removeArtikel(string judulArtikel, List_child &L, List_relasi &R);
void deleteLast(List_child &L, address_child &P);
void deleteAfter (List_child &L, address_child Prec, address_child &P);
#endif // CHILD_H_INCLUDED
