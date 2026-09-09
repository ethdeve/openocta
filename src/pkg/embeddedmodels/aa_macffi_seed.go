package embeddedmodels

// Seed Developer ID–signed libffi into the jupiterrider/ffi cache before yzma
// pulls in ffi's init (which dlopens libffi.8.dylib). File name sorts before
// other .go files so this blank import is registered first in the package.
import _ "github.com/openocta/openocta/pkg/macffi"
