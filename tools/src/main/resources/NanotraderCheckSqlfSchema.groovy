import groovy.sql.Sql
import org.apache.tools.ant.taskdefs.SQLExec
import org.apache.tools.ant.Project

def loadProps() {
  def props = new Properties()
  new File("nanotrader.sqlf.properties").withInputStream {
    stream -> props.load(stream)
  }
  p = new ConfigSlurper().parse(props)
}

def checkSchema() {
  Project project = new Project();
  project.init()
  project.setName("Nanotrader DB Check")
  
  //Check if nanotrader table exists
  def url = p.dbURLPrefix + p.dbHost + ":" + p.dbPort
  def sqlf
  for (int loop =0; loop < 120; loop++) {
    try {
      sqlf = Sql.newInstance(url, p.dbUser, p.dbPasswd, p.dbDriver)
    } catch(Exception e) {
      println "Error connecting to database: ${e.getMessage()}"
      sleep(1)
    }
  }
  def md = sqlf.connection.metaData

  for (int loop = 0; loop < 120; loop++) {
    try {
      rs = md.getTables(null, 'NANOTRADER', 'ACCOUNT' , null);
      account = rs.next();
    } catch(Exception e) {
      println "Error checking schema: ${e.getMessage()}"
    }

    if(account){
      println 'Schema exists'
      sqlf.close()
      System.exit(0)
      return
    }
    sleep(1)
  }

  println 'Schema not found'
  sqlf.close()
  System.exit(1)
}

loadProps()
checkSchema()
