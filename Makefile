include $(GOROOT)/src/Make.inc
 
TARG=gorel
GOFILES=gorel.go\
        nodes.go\
        visitor.go\
        to_sql.go\
        mysql.go\
include $(GOROOT)/src/Make.pkg 
