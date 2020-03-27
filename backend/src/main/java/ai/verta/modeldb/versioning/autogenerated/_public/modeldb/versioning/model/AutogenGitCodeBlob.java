// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model;

import ai.verta.modeldb.ModelDBException;
import ai.verta.modeldb.versioning.*;
import ai.verta.modeldb.versioning.blob.diff.*;
import ai.verta.modeldb.versioning.blob.diff.Function3;
import ai.verta.modeldb.versioning.blob.visitors.Visitor;
import com.pholser.junit.quickcheck.generator.*;
import com.pholser.junit.quickcheck.random.*;
import java.util.*;
import java.util.function.Function;

public class AutogenGitCodeBlob implements ProtoType {
  private String Branch;
  private String Hash;
  private Boolean IsDirty;
  private String Repo;
  private String Tag;

  public AutogenGitCodeBlob() {
    this.Branch = "";
    this.Hash = "";
    this.IsDirty = false;
    this.Repo = "";
    this.Tag = "";
  }

  public Boolean isEmpty() {
    if (this.Branch != null && !this.Branch.equals("")) {
      return false;
    }
    if (this.Hash != null && !this.Hash.equals("")) {
      return false;
    }
    if (this.IsDirty != null && !this.IsDirty.equals(false)) {
      return false;
    }
    if (this.Repo != null && !this.Repo.equals("")) {
      return false;
    }
    if (this.Tag != null && !this.Tag.equals("")) {
      return false;
    }
    return true;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("{\"class\": \"AutogenGitCodeBlob\", \"fields\": {");
    boolean first = true;
    if (this.Branch != null && !this.Branch.equals("")) {
      if (!first) sb.append(", ");
      sb.append("\"Branch\": " + "\"" + Branch + "\"");
      first = false;
    }
    if (this.Hash != null && !this.Hash.equals("")) {
      if (!first) sb.append(", ");
      sb.append("\"Hash\": " + "\"" + Hash + "\"");
      first = false;
    }
    if (this.IsDirty != null && !this.IsDirty.equals(false)) {
      if (!first) sb.append(", ");
      sb.append("\"IsDirty\": " + IsDirty);
      first = false;
    }
    if (this.Repo != null && !this.Repo.equals("")) {
      if (!first) sb.append(", ");
      sb.append("\"Repo\": " + "\"" + Repo + "\"");
      first = false;
    }
    if (this.Tag != null && !this.Tag.equals("")) {
      if (!first) sb.append(", ");
      sb.append("\"Tag\": " + "\"" + Tag + "\"");
      first = false;
    }
    sb.append("}}");
    return sb.toString();
  }

  // TODO: actually hash
  public String getSHA() {
    StringBuilder sb = new StringBuilder();
    sb.append("AutogenGitCodeBlob");
    if (this.Branch != null && !this.Branch.equals("")) {
      sb.append("::Branch::").append(Branch);
    }
    if (this.Hash != null && !this.Hash.equals("")) {
      sb.append("::Hash::").append(Hash);
    }
    if (this.IsDirty != null && !this.IsDirty.equals(false)) {
      sb.append("::IsDirty::").append(IsDirty);
    }
    if (this.Repo != null && !this.Repo.equals("")) {
      sb.append("::Repo::").append(Repo);
    }
    if (this.Tag != null && !this.Tag.equals("")) {
      sb.append("::Tag::").append(Tag);
    }

    return sb.toString();
  }

  // TODO: not consider order on lists
  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null) return false;
    if (!(o instanceof AutogenGitCodeBlob)) return false;
    AutogenGitCodeBlob other = (AutogenGitCodeBlob) o;

    {
      Function3<String, String, Boolean> f = (x, y) -> x.equals(y);
      if (this.Branch != null || other.Branch != null) {
        if (this.Branch == null && other.Branch != null) return false;
        if (this.Branch != null && other.Branch == null) return false;
        if (!f.apply(this.Branch, other.Branch)) return false;
      }
    }
    {
      Function3<String, String, Boolean> f = (x, y) -> x.equals(y);
      if (this.Hash != null || other.Hash != null) {
        if (this.Hash == null && other.Hash != null) return false;
        if (this.Hash != null && other.Hash == null) return false;
        if (!f.apply(this.Hash, other.Hash)) return false;
      }
    }
    {
      Function3<Boolean, Boolean, Boolean> f = (x, y) -> x.equals(y);
      if (this.IsDirty != null || other.IsDirty != null) {
        if (this.IsDirty == null && other.IsDirty != null) return false;
        if (this.IsDirty != null && other.IsDirty == null) return false;
        if (!f.apply(this.IsDirty, other.IsDirty)) return false;
      }
    }
    {
      Function3<String, String, Boolean> f = (x, y) -> x.equals(y);
      if (this.Repo != null || other.Repo != null) {
        if (this.Repo == null && other.Repo != null) return false;
        if (this.Repo != null && other.Repo == null) return false;
        if (!f.apply(this.Repo, other.Repo)) return false;
      }
    }
    {
      Function3<String, String, Boolean> f = (x, y) -> x.equals(y);
      if (this.Tag != null || other.Tag != null) {
        if (this.Tag == null && other.Tag != null) return false;
        if (this.Tag != null && other.Tag == null) return false;
        if (!f.apply(this.Tag, other.Tag)) return false;
      }
    }
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.Branch, this.Hash, this.IsDirty, this.Repo, this.Tag);
  }

  public AutogenGitCodeBlob setBranch(String value) {
    this.Branch = Utils.removeEmpty(value);
    return this;
  }

  public String getBranch() {
    return this.Branch;
  }

  public AutogenGitCodeBlob setHash(String value) {
    this.Hash = Utils.removeEmpty(value);
    return this;
  }

  public String getHash() {
    return this.Hash;
  }

  public AutogenGitCodeBlob setIsDirty(Boolean value) {
    this.IsDirty = Utils.removeEmpty(value);
    return this;
  }

  public Boolean getIsDirty() {
    return this.IsDirty;
  }

  public AutogenGitCodeBlob setRepo(String value) {
    this.Repo = Utils.removeEmpty(value);
    return this;
  }

  public String getRepo() {
    return this.Repo;
  }

  public AutogenGitCodeBlob setTag(String value) {
    this.Tag = Utils.removeEmpty(value);
    return this;
  }

  public String getTag() {
    return this.Tag;
  }

  public static AutogenGitCodeBlob fromProto(ai.verta.modeldb.versioning.GitCodeBlob blob) {
    if (blob == null) {
      return null;
    }

    AutogenGitCodeBlob obj = new AutogenGitCodeBlob();
    {
      Function<ai.verta.modeldb.versioning.GitCodeBlob, String> f = x -> (blob.getBranch());
      obj.Branch = Utils.removeEmpty(f.apply(blob));
    }
    {
      Function<ai.verta.modeldb.versioning.GitCodeBlob, String> f = x -> (blob.getHash());
      obj.Hash = Utils.removeEmpty(f.apply(blob));
    }
    {
      Function<ai.verta.modeldb.versioning.GitCodeBlob, Boolean> f = x -> (blob.getIsDirty());
      obj.IsDirty = Utils.removeEmpty(f.apply(blob));
    }
    {
      Function<ai.verta.modeldb.versioning.GitCodeBlob, String> f = x -> (blob.getRepo());
      obj.Repo = Utils.removeEmpty(f.apply(blob));
    }
    {
      Function<ai.verta.modeldb.versioning.GitCodeBlob, String> f = x -> (blob.getTag());
      obj.Tag = Utils.removeEmpty(f.apply(blob));
    }
    return obj;
  }

  public ai.verta.modeldb.versioning.GitCodeBlob.Builder toProto() {
    ai.verta.modeldb.versioning.GitCodeBlob.Builder builder =
        ai.verta.modeldb.versioning.GitCodeBlob.newBuilder();
    {
      if (this.Branch != null && !this.Branch.equals("")) {
        Function<ai.verta.modeldb.versioning.GitCodeBlob.Builder, Void> f =
            x -> {
              builder.setBranch(this.Branch);
              return null;
            };
        f.apply(builder);
      }
    }
    {
      if (this.Hash != null && !this.Hash.equals("")) {
        Function<ai.verta.modeldb.versioning.GitCodeBlob.Builder, Void> f =
            x -> {
              builder.setHash(this.Hash);
              return null;
            };
        f.apply(builder);
      }
    }
    {
      if (this.IsDirty != null && !this.IsDirty.equals(false)) {
        Function<ai.verta.modeldb.versioning.GitCodeBlob.Builder, Void> f =
            x -> {
              builder.setIsDirty(this.IsDirty);
              return null;
            };
        f.apply(builder);
      }
    }
    {
      if (this.Repo != null && !this.Repo.equals("")) {
        Function<ai.verta.modeldb.versioning.GitCodeBlob.Builder, Void> f =
            x -> {
              builder.setRepo(this.Repo);
              return null;
            };
        f.apply(builder);
      }
    }
    {
      if (this.Tag != null && !this.Tag.equals("")) {
        Function<ai.verta.modeldb.versioning.GitCodeBlob.Builder, Void> f =
            x -> {
              builder.setTag(this.Tag);
              return null;
            };
        f.apply(builder);
      }
    }
    return builder;
  }

  public void preVisitShallow(Visitor visitor) throws ModelDBException {
    visitor.preVisitAutogenGitCodeBlob(this);
  }

  public void preVisitDeep(Visitor visitor) throws ModelDBException {
    this.preVisitShallow(visitor);
    visitor.preVisitDeepString(this.Branch);
    visitor.preVisitDeepString(this.Hash);
    visitor.preVisitDeepBoolean(this.IsDirty);
    visitor.preVisitDeepString(this.Repo);
    visitor.preVisitDeepString(this.Tag);
  }

  public AutogenGitCodeBlob postVisitShallow(Visitor visitor) throws ModelDBException {
    return visitor.postVisitAutogenGitCodeBlob(this);
  }

  public AutogenGitCodeBlob postVisitDeep(Visitor visitor) throws ModelDBException {
    this.setBranch(visitor.postVisitDeepString(this.Branch));
    this.setHash(visitor.postVisitDeepString(this.Hash));
    this.setIsDirty(visitor.postVisitDeepBoolean(this.IsDirty));
    this.setRepo(visitor.postVisitDeepString(this.Repo));
    this.setTag(visitor.postVisitDeepString(this.Tag));
    return this.postVisitShallow(visitor);
  }
}