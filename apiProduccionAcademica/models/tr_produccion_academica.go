package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrProduccionAcademica struct {
	ProduccionAcademica *ProduccionAcademica
	DatoProduccion      *[]DatoProduccion
}

//funcion para la transaccion de ProduccionAcademica -Agregar una ProduccionAcademica
func AddTransaccionProduccionAcademica(m *TrProduccionAcademica) (id int64, err error) {
	o := orm.NewOrm()
	fmt.Println("ProduccionAcademica!!!!", m.ProduccionAcademica.Id)
	o.Begin()

	if id, err := o.Insert(m.ProduccionAcademica); err == nil {
		m.ProduccionAcademica.Id = int(id)

		fmt.Println("ProduccionAcademica", m.ProduccionAcademica)
		for _, v := range *m.DatoProduccion {
			v.ProduccionAcademicaId = m.ProduccionAcademica
			if _, err = o.Insert(&v); err != nil {
				fmt.Println("DatoProduccion", &v)
				err = o.Rollback()
			}
			fmt.Println("DatoProduccion", &v)
		}

		err = o.Commit()
	} else {
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
