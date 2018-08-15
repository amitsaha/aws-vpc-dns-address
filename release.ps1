$VERSION=$(cat VERSION)
$DISTDIR="dist/$VERSION"
mkdir -p $DISTDIR
git tag --annotate v$VERSION -m "Release v$VERSION"

$pairs = @('linux/amd64', 'windows/amd64')
foreach ($pair in $pairs)
{
	GOOS=`echo $pair | cut -d'/' -f1`
    GOARCH=`echo $pair | cut -d'/' -f2` 
    OBJECT_FILE="aws-vpc-dns-address-$VERSION-$GOOS-$GOARCH"
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$DISTDIR/$OBJECT_FILE" 
    Push-Location $DISTDIR
    echo $OBJECT_FILE
    Compress-Archive -Path $OBJECT_FILE -DestinationPath "$OBJECT_FILE.zip"
    Pop-Location
done
git push --tags