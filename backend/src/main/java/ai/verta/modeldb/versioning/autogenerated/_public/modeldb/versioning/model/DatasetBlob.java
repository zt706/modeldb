// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model;

import java.util.*;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import ai.verta.modeldb.ModelDBException;
import ai.verta.modeldb.versioning.*;
import ai.verta.modeldb.versioning.blob.diff.Function3;
import ai.verta.modeldb.versioning.blob.diff.*;
import ai.verta.modeldb.versioning.blob.visitors.Visitor;
import com.pholser.junit.quickcheck.generator.*;
import com.pholser.junit.quickcheck.random.*;

public class DatasetBlob implements ProtoType {
    public S3DatasetBlob S3;
    public PathDatasetBlob Path;

    public DatasetBlob() {
        this.S3 = null;
        this.Path = null;
    }

    public Boolean isEmpty() {
        if (this.S3 != null && !this.S3.equals(null) ) {
            return false;
        }
        if (this.Path != null && !this.Path.equals(null) ) {
            return false;
        }
        return true;
    }

    @Override
    public String toString() {
        return "{\"class\": \"DatasetBlob\",\"fields\": {" +
                "\"S3\": " + S3 + ", " +
                "\"Path\": " + Path + 
                "}}";
    }

    // TODO: not consider order on lists
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null) return false;
        if (!(o instanceof DatasetBlob)) return false;
        DatasetBlob other = (DatasetBlob) o;

        {
            Function3<S3DatasetBlob,S3DatasetBlob,Boolean> f = (x, y) -> x.equals(y);
            if (this.S3 != null || other.S3 != null) {
                if (this.S3 == null && other.S3 != null)
                    return false;
                if (this.S3 != null && other.S3 == null)
                    return false;
                if (!f.apply(this.S3, other.S3))
                    return false;
            }
        }
        {
            Function3<PathDatasetBlob,PathDatasetBlob,Boolean> f = (x, y) -> x.equals(y);
            if (this.Path != null || other.Path != null) {
                if (this.Path == null && other.Path != null)
                    return false;
                if (this.Path != null && other.Path == null)
                    return false;
                if (!f.apply(this.Path, other.Path))
                    return false;
            }
        }
        return true;
    }

    @Override
    public int hashCode() {
        return Objects.hash(
        this.S3,
        this.Path
        );
      }

    public DatasetBlob setS3(S3DatasetBlob value) {
        this.S3 = Utils.removeEmpty(value);
        return this;
    }
    public DatasetBlob setPath(PathDatasetBlob value) {
        this.Path = Utils.removeEmpty(value);
        return this;
    }

    static public DatasetBlob fromProto(ai.verta.modeldb.versioning.DatasetBlob blob) {
        if (blob == null) {
            return null;
        }

        DatasetBlob obj = new DatasetBlob();
        {
            Function<ai.verta.modeldb.versioning.DatasetBlob,S3DatasetBlob> f = x -> S3DatasetBlob.fromProto(blob.getS3());
            obj.S3 = Utils.removeEmpty(f.apply(blob));
        }
        {
            Function<ai.verta.modeldb.versioning.DatasetBlob,PathDatasetBlob> f = x -> PathDatasetBlob.fromProto(blob.getPath());
            obj.Path = Utils.removeEmpty(f.apply(blob));
        }
        return obj;
    }

    public ai.verta.modeldb.versioning.DatasetBlob.Builder toProto() {
        ai.verta.modeldb.versioning.DatasetBlob.Builder builder = ai.verta.modeldb.versioning.DatasetBlob.newBuilder();
        {
            if (this.S3 != null && !this.S3.equals(null) ) {
                Function<ai.verta.modeldb.versioning.DatasetBlob.Builder,Void> f = x -> { builder.setS3(this.S3.toProto()); return null; };
                f.apply(builder);
            }
        }
        {
            if (this.Path != null && !this.Path.equals(null) ) {
                Function<ai.verta.modeldb.versioning.DatasetBlob.Builder,Void> f = x -> { builder.setPath(this.Path.toProto()); return null; };
                f.apply(builder);
            }
        }
        return builder;
    }

    public void preVisitShallow(Visitor visitor) throws ModelDBException {
        visitor.preVisitDatasetBlob(this);
    }

    public void preVisitDeep(Visitor visitor) throws ModelDBException {
        this.preVisitShallow(visitor);
        visitor.preVisitDeepS3DatasetBlob(this.S3);
        visitor.preVisitDeepPathDatasetBlob(this.Path);
    }

    public DatasetBlob postVisitShallow(Visitor visitor) throws ModelDBException {
        return visitor.postVisitDatasetBlob(this);
    }

    public DatasetBlob postVisitDeep(Visitor visitor) throws ModelDBException {
        this.setS3(visitor.postVisitDeepS3DatasetBlob(this.S3));
        this.setPath(visitor.postVisitDeepPathDatasetBlob(this.Path));
        return this.postVisitShallow(visitor);
    }
}
