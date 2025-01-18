#ifndef PARENT_H_INCLUDED
#define PARENT_H_INCLUDED
#include <iostream>
#define next(P) P->next
#define prev(P) P->prev
#define first(L) L.first
#define info(P) P->info
#define child(P) P->child
#define parent(P) P->parent
#include "child.h"
#include "relasi.h"

using namespace std;

struct List_child;
struct List_relasi;
typedef struct elmlist_parent *address_parent;
struct penulis {
    string nama;
    string id;
    string email;
};
typedef penulis infotype_parent;
struct elmlist_parent {
    infotype_parent info;
    address_parent next;

};
struct List_parent {
    address_parent first;
};

void createList(List_parent &L);
void insertFirst(List_parent &L, address_parent P);
void insertAfter(List_parent &L, address_parent Prec, address_parent P);
void insertLast(List_parent &L, address_parent P);
void deleteFirst(List_parent &L, address_parent &P);
void deleteLast(List_parent &L, address_parent &P);
void deleteAfter(List_parent &L, address_parent Prec, address_parent &P);
address_parent alokasiParent(infotype_parent x);
void dealokasi(address_parent &P);
address_parent findElm_nama(List_parent L, string x);
void searchElm_nama(List_parent L, string x);
address_parent findElm_id(List_parent L, string x);
void searchElm_id(List_parent L, string x);
void searchElm_email(List_parent L, string x);
void printInfo(List_parent L);
void removePenulis(string nama, List_parent &L, List_child &C, List_relasi &R);

#endif // PARENT_H_INCLUDED
