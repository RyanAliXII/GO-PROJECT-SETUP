package web

import (
	"fmt"
	"ryanali12/web_service/web/routes"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	fmt.Println(`                                                                                                                                                                     
                                                                                                                                                                     
	TTTTTTTTTTTTTTTTTTTTTTTEEEEEEEEEEEEEEEEEEEEEE   SSSSSSSSSSSSSSS TTTTTTTTTTTTTTTTTTTTTTT                    AAA               PPPPPPPPPPPPPPPPP   PPPPPPPPPPPPPPPPP   
	T:::::::::::::::::::::TE::::::::::::::::::::E SS:::::::::::::::ST:::::::::::::::::::::T                   A:::A              P::::::::::::::::P  P::::::::::::::::P  
	T:::::::::::::::::::::TE::::::::::::::::::::ES:::::SSSSSS::::::ST:::::::::::::::::::::T                  A:::::A             P::::::PPPPPP:::::P P::::::PPPPPP:::::P 
	T:::::TT:::::::TT:::::TEE::::::EEEEEEEEE::::ES:::::S     SSSSSSST:::::TT:::::::TT:::::T                 A:::::::A            PP:::::P     P:::::PPP:::::P     P:::::P
	TTTTTT  T:::::T  TTTTTT  E:::::E       EEEEEES:::::S            TTTTTT  T:::::T  TTTTTT                A:::::::::A             P::::P     P:::::P  P::::P     P:::::P
			T:::::T          E:::::E             S:::::S                    T:::::T                       A:::::A:::::A            P::::P     P:::::P  P::::P     P:::::P
			T:::::T          E::::::EEEEEEEEEE    S::::SSSS                 T:::::T                      A:::::A A:::::A           P::::PPPPPP:::::P   P::::PPPPPP:::::P 
			T:::::T          E:::::::::::::::E     SS::::::SSSSS            T:::::T                     A:::::A   A:::::A          P:::::::::::::PP    P:::::::::::::PP  
			T:::::T          E:::::::::::::::E       SSS::::::::SS          T:::::T                    A:::::A     A:::::A         P::::PPPPPPPPP      P::::PPPPPPPPP    
			T:::::T          E::::::EEEEEEEEEE          SSSSSS::::S         T:::::T                   A:::::AAAAAAAAA:::::A        P::::P              P::::P            
			T:::::T          E:::::E                         S:::::S        T:::::T                  A:::::::::::::::::::::A       P::::P              P::::P            
			T:::::T          E:::::E       EEEEEE            S:::::S        T:::::T                 A:::::AAAAAAAAAAAAA:::::A      P::::P              P::::P            
		  TT:::::::TT      EE::::::EEEEEEEE:::::ESSSSSSS     S:::::S      TT:::::::TT              A:::::A             A:::::A   PP::::::PP          PP::::::PP          
		  T:::::::::T      E::::::::::::::::::::ES::::::SSSSSS:::::S      T:::::::::T             A:::::A               A:::::A  P::::::::P          P::::::::P          
		  T:::::::::T      E::::::::::::::::::::ES:::::::::::::::SS       T:::::::::T            A:::::A                 A:::::A P::::::::P          P::::::::P          
		  TTTTTTTTTTT      EEEEEEEEEEEEEEEEEEEEEE SSSSSSSSSSSSSSS         TTTTTTTTTTT           AAAAAAA                   AAAAAAAPPPPPPPPPP          PPPPPPPPPP  `)

	r.GET("/home", func(ctx * gin.Context){
		ctx.HTML(http.StatusOK, "pages/home.html", nil)
	})
	routes.InitProductRoutes(r)
}
