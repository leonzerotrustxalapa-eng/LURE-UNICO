#include <stdio.h>
#include <time.h>

#define LURE_TIER "TIER_C4"
#define LURE_SELLO "LURE-LOGv1 | LEON-XALAPA | TIER_C4 | CDA2A246C493955D"

void registrar_evento_c4() {
    time_t t = time(NULL);
    printf("[%s] [LURE-C] Evento C4 blindado registrado - %s", LURE_TIER, ctime(&t));
    printf("%s\n", LURE_SELLO);
}

int main() {
    printf("=== NIVEL 1 C4 CORE - C ===\n");
    registrar_evento_c4();
    printf("CERO ERRORES | BLINDADO\n");
    return 0;
}
