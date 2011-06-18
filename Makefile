include $(GOROOT)/src/Make.inc
 
TARG=gorel
GOFILES=gorel.go\
        nodes/nodes.go\
        visitors/visitor.go\
        visitors/to_sql.go
include $(GOROOT)/src/Make.pkg 
