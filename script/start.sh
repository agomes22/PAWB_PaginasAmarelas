PROJECT="paginas"
base="/go/src/"
path=$base$PROJECT
pwd
echo $path
if [ ! -d $path ] 
then
	revel new -a $PROJECT
fi
cd $path
revel run -a $PROJECT
