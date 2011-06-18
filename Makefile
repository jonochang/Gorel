include $(GOROOT)/src/Make.inc
 
TARG=gorel
GOFILES=gorel.go\
        generated/nodes.go\
        generated/visitor.go\
        to_sql.go
include $(GOROOT)/src/Make.pkg 
