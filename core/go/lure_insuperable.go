package main
import ("fmt"; "time"; "crypto/sha256")
const SELLO = "LURE-LOGv1 | LEON-XALAPA | TIER_C4 | CDA2A246C493955D"
func main(){
  h := sha256.Sum256([]byte(SELLO))
  fmt.Println("\n[TIER_C4 - GO] NUCLEO NUBE MUNDIAL")
  fmt.Println("Sello:", SELLO)
  fmt.Printf("SHA256: %x\n", h)
  fmt.Println("ESTADO: CERO ERRORES - INSUPERABLE")
  fmt.Println("TODO TIENE MI PROGRAMA |", time.Now().Format("2006-01-02 15:04:05"))
}
