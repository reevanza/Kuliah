#ifndef RELASI_H_INCLUDED
#define RELASI_H_INCLUDED
#include <iostream>
#define next(P) P->next
#define prev(P) P->prev
#define first(L) L.first
#define last(L) L.last
#define info(P) P->info
#define child(P) P->child
#define parent(P) P->parent
using namespace std;
#include "child.h"
#include "parent.h"

struct elmlist_child;
struct elmlist_parent;
typedef struct elmlist_child *address_child;
typedef struct elmlist_parent *address_parent;
typedef struct elmlist_relasi *address_relasi;

struct elmlist_relasi{
    address_relasi next;
    address_child child;
    address_parent parent;
};

struct List_relasi{
    address_relasi first;

};
void createList(List_relasi &L);
void insertFirst(List_relasi &L, address_relasi P);
address_relasi alokasiRelasi( address_parent P, address_child C);
void deleteFirst(List_relasi &L, address_relasi &P);
void dealokasi(address_relasi &P);
address_relasi findElm(List_relasi L, address_parent P, address_child C);
void printInfo(List_relasi L);
void deleteRelasiParent(List_relasi &L, address_parent P);
void deleteRelasiChild(List_relasi &R, address_child P);
void countArtikel(List_relasi R, address_parent P);
void deleteRelasiParentChild(List_relasi &R, string namaParent, string judulArtikel);
void searchArtikelberdasarkanParent(List_relasi &R, string namaParent, string judulArtikel);

#endif // RELASI_H_INCLUDED
