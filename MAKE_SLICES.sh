#!/bin/sh

#all go types
gotypes="string,int,uint,uint8,uint16,uint32,uint64"
gotypes="$gotypes,int8,int16,int32,int64,float32,float64"
gotypes="$gotypes,complex64,complex128,byte,rune"

#may add your types
mytypes=""

gotypes="$gotypes,$mytypes"

#delete everything after this line:
sed -n -i '/\/\/ ----------delete everything below this line --------------- \/\//q;p' slices.go
echo "// ----------delete everything below this line --------------- //" >> slices.go
sed -n -i '/\/\/ ----------delete everything below this line --------------- \/\//q;p' channel.go
echo "// ----------delete everything below this line --------------- //" >> channel.go

#read slices.go.raw and channel.go.raw
#replace #TYPE and #TYPECAMEL accordingly
for type in $(echo $gotypes | tr "," "\n")
do
	typecamel=$(echo ${type:0:1} | tr  '[a-z]' '[A-Z]')${type:1}
	sed "s/#TYPECAMEL/$typecamel/g;s/#TYPE/$type/g" slices.go.raw >> slices.go
	sed "s/#TYPECAMEL/$typecamel/g;s/#TYPE/$type/g" channel.go.raw >> channel.go
done
