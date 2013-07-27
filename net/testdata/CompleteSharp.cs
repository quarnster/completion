/*
Copyright (c) 2012 Fredrik Ehnbom

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

   1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.

   2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.

   3. This notice may not be removed or altered from any source
   distribution.
*/
using System;
using System.Reflection;
using System.Collections;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.IO;

public class CompleteSharp
{
    private static string sep = ";;--;;";

    protected class MyAppDomain : MarshalByRefObject
    {
        class Hack : MarshalByRefObject
        {
            public Hack()
            {
                AppDomain.CurrentDomain.AssemblyResolve += ResolveAssembly;
            }

            private int GetParameterExtent(string parameter)
            {
                int hookcount = 0;
                int ltgtcount = 0;
                for (int i = 0; i < parameter.Length; i++)
                {
                    switch (parameter[i])
                    {
                        case ']':
                        {
                            hookcount--;
                            if (hookcount == 0 && ltgtcount == 0)
                            {
                                return i+1;
                            }
                            break;
                        }
                        case '[': hookcount++; break;
                        case ',':
                        {
                            if (parameter[i] == ',' && hookcount == 0 && ltgtcount == 0)
                            {
                                return i;
                            }
                            break;
                        }
                        case '<': ltgtcount++; break;
                        case '>':
                        {
                            ltgtcount--;
                            if (hookcount == 0 && ltgtcount == 0)
                            {
                                return i+1;
                            }
                            break;
                        }
                    }
                }
                return parameter.Length;
            }
            private string[] SplitParameters(string parameters, bool fix=true)
            {
                List<string> s = new List<string>();
                for (int i = 0; i < parameters.Length;)
                {
                    int len = GetParameterExtent(parameters.Substring(i));
                    string toadd = parameters.Substring(i, len);
                    while (toadd.Length >= 2 && toadd.StartsWith("["))
                    {
                        toadd = toadd.Substring(1, toadd.Length-2);
                        toadd = toadd.Substring(0, GetParameterExtent(toadd));
                    }
                    if (fix && toadd.Trim().Length > 0)
                        toadd = FixName(toadd);
                    toadd = toadd.Trim();
                    if (toadd.Length > 0)
                        s.Add(toadd);
                    i += len;
                    if (len == 0)
                        i++;
                }
                return s.ToArray();
            }

            private string ParseParameters(string parameters, int expected, bool insertion)
            {
                if (parameters.Length >= 2 && parameters.StartsWith("["))
                {
                    parameters = parameters.Substring(1, parameters.Length-2);
                }
                string[] para = null;
                if (parameters.Length > 0)
                {
                    para = SplitParameters(parameters);
                }
                else
                {
                    para = new string[expected];
                    for (int i = 0; i < expected; i++)
                    {
                        para[i] = "T" + (i+1);
                    }
                }
                string ret = "";
                for (int i = 0; i < para.Length; i++)
                {
                    if (ret.Length > 0)
                        ret += ", ";
                    if (!insertion)
                    {
                        ret+= para[i];
                    }
                    else
                    {
                        ret += "${" + (i+1) + ":" + para[i] + "}";
                    }
                }
                return ret;
            }
            private string FixName(string str, bool insertion=false)
            {
                int index = str.IndexOf('`');
                if (index != -1)
                {
                    Regex regex = new Regex("^\\s*([^`]+)\\`(\\d+)([^\\[]+)?(\\[.*\\])?");
                    Match m = regex.Match(str.Replace("$", "\\$"));
                    string type = m.Groups[1].ToString();
                    int num = System.Int32.Parse(m.Groups[2].ToString());
                    string subclass = m.Groups[3].ToString();
                    string parameters = m.Groups[4].ToString();
                    if (subclass.Length > 0)
                    {
                        subclass = "." + subclass.Substring(1);
                    }
                    string extra = "";
                    while (parameters.EndsWith("[]"))
                    {
                        extra += "[]";
                        parameters = parameters.Substring(0, parameters.Length-2);
                    }


                    str = type + "<" + ParseParameters(parameters, num, insertion) + ">" + subclass + extra;
                }
                return str;
            }

            private string[] GetTemplateArguments(string template)
            {
                int index = template.IndexOf('<');
                int index2 = template.LastIndexOf('>');
                if (index != -1 && index2 != -1)
                {
                    string args = template.Substring(index+1, index2-index-1);
                    return SplitParameters(args, false);
                }
                return new string[0];
            }

            private string GetBase(string fullname)
            {
                int index = fullname.IndexOf('<');
                if (index == -1)
                    return fullname;
                return fullname.Substring(0, index);
            }

            protected Type GetType(MyAppDomain ad, string basename, string[] templateParam)
            {
                if (templateParam.Length > 0 && basename.IndexOf('`') == -1)
                {
                    basename += "`" + templateParam.Length;
                }
                Type[] subtypes = new Type[templateParam.Length];
                for (int i = 0; i < subtypes.Length; i++)
                {
                    string bn = GetBase(templateParam[i]);
                    string[] args = GetTemplateArguments(templateParam[i]);
                    subtypes[i] = GetType(ad, bn, args);
                }

                Type t = GetType(basename);

                if (t != null && subtypes.Length > 0)
                {
                    try
                    {
                        Type t2 = t.MakeGenericType(subtypes);
                        System.Console.Error.WriteLine("returning type2: " + t2.FullName);
                        return t2;
                    }
                    catch (Exception e)
                    {
                        reportError(e);
                    }
                }
                if (t != null)
                    System.Console.Error.WriteLine("returning type: " + t.FullName);
                return t;
            }
            private enum Accessibility
            {
                NONE = (0<<0),
                STATIC = (1<<0),
                PRIVATE = (1<<1),
                PROTECTED = (1<<2),
                PUBLIC = (1<<3),
                INTERNAL = (1<<4)
            };

            private int GetModifiers(MemberInfo m)
            {
                Accessibility modifiers = Accessibility.NONE;
                switch (m.MemberType)
                {
                    case MemberTypes.Field:
                    {
                        FieldInfo f = (FieldInfo)m;
                        if (f.IsPrivate)
                            modifiers |= Accessibility.PRIVATE;
                        if (f.IsPublic)
                            modifiers |= Accessibility.PUBLIC;
                        if (f.IsStatic)
                            modifiers |= Accessibility.STATIC;
                        if (!f.IsPublic && !f.IsPrivate)
                            modifiers |= Accessibility.PROTECTED;
                        break;
                    }
                    case MemberTypes.Method:
                    {
                        MethodInfo mi = (MethodInfo)m;
                        if (mi.IsPrivate)
                            modifiers |= Accessibility.PRIVATE;
                        if (mi.IsPublic)
                            modifiers |= Accessibility.PUBLIC;
                        if (mi.IsStatic)
                            modifiers |= Accessibility.STATIC;
                        if (!mi.IsPublic && !mi.IsPrivate)
                            modifiers |= Accessibility.PROTECTED;
                        break;
                    }
                    case MemberTypes.Property:
                    {
                        PropertyInfo p = (PropertyInfo)m;
                        foreach (MethodInfo mi in p.GetAccessors())
                        {
                            modifiers |= (Accessibility)GetModifiers(mi);
                        }
                        break;
                    }
                    default:
                    {
                        modifiers = Accessibility.STATIC|Accessibility.PUBLIC;
                        break;
                    }
                }
                return (int) modifiers;
            }

            public MyAppDomain ad;


            Assembly TryLoad(string path, string name)
            {
                path = path.Replace("file://", "");
                int idx1 = path.LastIndexOf('/');
                int idx2 = path.LastIndexOf('\\');
                int idx = Math.Max(idx1, idx2);
                if (idx != -1)
                {
                    path = path.Substring(0, idx);
                }
                path = path + "/" + name + ".dll";
                FileInfo fi = new FileInfo(path);
                if (fi.Exists)
                {
                    return Load(path);
                }
                return null;
            }
            private Assembly ResolveAssembly(object sender, ResolveEventArgs args)
            {
                string name = args.Name.Substring(0, args.Name.IndexOf(","));
                foreach (Assembly asm in AppDomain.CurrentDomain.GetAssemblies())
                {
                    if (asm.FullName.Substring(0, asm.FullName.IndexOf(',')) == name)
                    {
                        return asm;
                    }
                }
                // Assembly wasn't found, try and see if it's in the same directory
                // as one of the other already loaded assemblies.
                if (ad != null)
                {
                    foreach (string asm in ad.assemblies)
                    {
                        Assembly ret = TryLoad(asm, name);
                        if (ret != null)
                            return ret;
                    }
                }

                foreach (Assembly asm in AppDomain.CurrentDomain.GetAssemblies())
                {
                    Assembly ret = TryLoad(asm.CodeBase, name);
                    if (ret != null)
                        return ret;
                }
                System.Console.Error.WriteLine("Couldn't find assembly: " + name);
                return null;
            }
            public Assembly Load(String str)
            {
                return Assembly.Load(File.ReadAllBytes(str));
            }
            public Type GetType(string basename)
            {
                Type t = null;
                foreach (Assembly ass in AppDomain.CurrentDomain.GetAssemblies())
                {
                    t = ass.GetType(basename);
                    if (t != null)
                        break;
                }
                return t;
            }
            public bool Execute(string[] args, ArrayList modules)
            {
                try
                {
                    foreach (string s in args)
                    {
                        System.Console.Error.WriteLine("execute args: " + s);
                    }
                    Assembly[] assemblies = AppDomain.CurrentDomain.GetAssemblies();
                    if (args[0] == "-quit")
                    {
                        return false;
                    }
                    else if (args[0] == "-findclass")
                    {
                        bool found = false;
                        foreach (String mod in modules)
                        {
                            try
                            {
                                string fullname = args[1];
                                if (mod.Length > 0)
                                {
                                    fullname = mod + "." + fullname;
                                }
                                System.Console.Error.WriteLine("Trying " + fullname);
                                Type t2 = GetType(fullname);
                                if (t2 != null)
                                {
                                    string typename = t2.FullName;
                                    System.Console.WriteLine(typename);
                                    System.Console.Error.WriteLine(typename);
                                    found = true;
                                    break;
                                }
                            }
                            catch (Exception e)
                            {
                                reportError(e);
                            }
                        }
                        if (found)
                            return true;
                        // Probably a namespace then?

                        foreach (Assembly asm in assemblies)
                        {
                            try
                            {
                                foreach (Type t3 in asm.GetTypes())
                                {
                                    if (t3.Namespace == null)
                                        continue;
                                    foreach (string mod in modules)
                                    {
                                        string test = args[1];
                                        if (mod.Length > 0)
                                            test = mod + "." + test;
                                        if (t3.Namespace.StartsWith(test))
                                        {
                                            System.Console.WriteLine(test);
                                            System.Console.Error.WriteLine(test);
                                            found = true;
                                            break;
                                        }
                                    }
                                    if (found)
                                        break;
                                }
                                if (found)
                                    break;
                            }
                            catch (Exception)
                            {
                            }
                        }
                        return true;
                    }
                    if (args.Length < 2)
                        return true;
                        int len = args.Length - 3;
                    if (len < 0)
                        len = 0;
                    string[] templateParam = new string[len];
                    for (int i = 0; i < len; i++)
                    {
                        templateParam[i] = args[i+3];
                    }
                    Type t = null;
                    try
                    {
                        t = GetType(ad, args[1], templateParam);
                    }
                    catch (Exception e)
                    {
                        reportError(e);
                    }

                    if (args[0] == "-complete")
                    {
                        if (t != null)
                        {
                            foreach (MemberInfo m in t.GetMembers(BindingFlags.Instance|BindingFlags.Static|BindingFlags.Public|BindingFlags.NonPublic|BindingFlags.FlattenHierarchy))
                            {
                                switch (m.MemberType)
                                {
                                    case MemberTypes.Event:
                                    case MemberTypes.Field:
                                    case MemberTypes.Method:
                                    case MemberTypes.Property:
                                    {
                                        string completion = null;
                                        try
                                        {
                                            completion = m.ToString();
                                        }
                                        catch (Exception e)
                                        {
                                            System.Console.Error.WriteLine("Skipping this member as an exception was thrown when calling ToString:");
                                            System.Console.Error.WriteLine(e.Message);
                                            // Seems that some state is messed up when this happens, so we need to
                                            // force a reload of the assemblies to make it work again.
                                            ad.forceReload = true;
                                            continue;
                                        }
                                        int index = completion.IndexOf(' ');
                                        string returnType = completion.Substring(0, index);
                                        completion = completion.Substring(index+1);

                                        string display = "";
                                        index = completion.IndexOf('(');
                                        int index2 = completion.LastIndexOf(')');
                                        if (index != -1 && index2 != -1)
                                        {
                                            string param = completion.Substring(index+1, index2-index-1);
                                            completion = completion.Substring(0, index+1);
                                            display = completion;
                                            string[] par = param.Split(new Char[]{','});
                                            int addIndex = 1;
                                            index2 = completion[index-1] == ']' ? index-1 : -1;
                                            display = completion;
                                            bool first = true;
                                            string[] generics = null;
                                            if (index2 != -1)
                                            {
                                                index = completion.IndexOf('[')+1;
                                                // Generic method
                                                string generic = completion.Substring(index, index2-index);
                                                completion = completion.Substring(0, index-1) + "<";
                                                display = completion;
                                                generics = generic.Split(new char[]{','});
                                                foreach (string g in generics)
                                                {
                                                    if (!first)
                                                    {
                                                        completion += ", ";
                                                        display += ", ";
                                                    }
                                                    display += g;
                                                    completion += "${" + addIndex + ":" + g + "}";
                                                    addIndex++;
                                                    first = false;
                                                }
                                                display += ">(";
                                                completion += ">(";
                                            }
                                            first = true;
                                            foreach (string p in par)
                                            {
                                                string toadd = FixName(p.Trim());
                                                if (toadd.Length > 0)
                                                {
                                                    if (!first)
                                                    {
                                                        completion += ", ";
                                                        display += ", ";
                                                    }
                                                    display += toadd;
                                                    completion += "${" + addIndex + ":" + toadd + "}";
                                                    addIndex++;
                                                    first = false;
                                                }
                                            }
                                            completion += ")";
                                            display += ")";
                                        }
                                        else
                                        {
                                            display = completion;
                                        }
                                        string insertion = completion;
                                        display += "\t" + FixName(returnType);

                                        System.Console.WriteLine(display + sep + insertion + sep + GetModifiers(m));
                                        break;
                                    }
                                }
                            }
                        }
                        else
                        {
                            foreach (Assembly asm in assemblies)
                            {
                                try
                                {
                                    foreach (Type t3 in asm.GetTypes())
                                    {
                                        if (t3.Namespace == null)
                                            continue;
                                        if (t3.Namespace == args[1])
                                        {
                                            System.Console.WriteLine(FixName(t3.Name) + "\tclass" + sep + FixName(t3.Name, true));
                                        }
                                        else if (t3.Namespace != args[1] && t3.Namespace.StartsWith(args[1]))
                                        {
                                            string name = t3.Namespace.Substring(args[1].Length+1);
                                            int index = name.IndexOf('.');
                                            if (index != -1)
                                                name = name.Substring(0, index);

                                            System.Console.WriteLine(name + "\tnamespace" + sep + name);
                                        }
                                    }
                                }
                                catch (Exception)
                                {
                                }
                            }
                        }
                    }
                    else if (args[0] == "-returntype")
                    {
                        if (t != null)
                        {
                            bool found = false;
                            BindingFlags flags = BindingFlags.Instance|BindingFlags.Static|BindingFlags.Public|BindingFlags.NonPublic|BindingFlags.FlattenHierarchy;
                            string funcname = args[2];
                            Type[] generics = null;
                            int index = funcname.IndexOf('<');
                            int index2 = funcname.LastIndexOf('>');
                            if (index != -1 && index2 != -1)
                            {
                                string[] generics2 = funcname.Substring(index+1, index2-index-1).Split(new Char[]{','});
                                funcname = funcname.Substring(0, index);
                                generics = new Type[generics2.Length];
                                for (int i = 0; i < generics2.Length; i++)
                                {
                                    generics[i] = GetType(generics2[i]);
                                    if (generics[i] == null)
                                    {
                                        generics = null;
                                        break;
                                    }
                                }
                            }
                            // This isn't 100% correct, but an instance where two things
                            // are named the same but return two different types would
                            // be considered rare.
                            foreach (MethodInfo m in t.GetMethods(flags))
                            {
                                if (m.Name == funcname)
                                {
                                    MethodInfo m2 = m;
                                    if (generics != null)
                                    {
                                        m2 = m2.MakeGenericMethod(generics);
                                    }
                                    System.Console.WriteLine(FixName(m2.ReturnType.FullName));
                                    found = true;
                                    break;
                                }
                            }
                            if (found)
                                return true;
                            foreach (FieldInfo f in t.GetFields(flags))
                            {
                                if (f.Name == args[2])
                                {
                                    System.Console.WriteLine(FixName(f.FieldType.FullName));
                                    found = true;
                                    break;
                                }
                            }
                            if (found)
                                return true;
                            foreach (EventInfo e in t.GetEvents(flags))
                            {
                                if (e.Name == args[2])
                                {
                                    System.Console.WriteLine(FixName(e.EventHandlerType.FullName));
                                    found = true;
                                    break;
                                }
                            }
                            if (found)
                                return true;
                            foreach (PropertyInfo p in t.GetProperties(flags))
                            {
                                if (p.Name == args[2])
                                {
                                    System.Console.WriteLine(FixName(p.PropertyType.FullName));
                                    found = true;
                                    break;
                                }
                            }
                        }
                        else
                        {
                            bool found = false;
                            foreach (Assembly asm in assemblies)
                            {
                                try
                                {
                                    foreach (Type t3 in asm.GetTypes())
                                    {
                                        if (t3.Namespace == null)
                                            continue;
                                        if (t3.Namespace == args[1] && t3.Name == args[2])
                                        {
                                            System.Console.WriteLine(FixName(t3.FullName));
                                            found = true;
                                            break;
                                        }
                                    }
                                    if (found)
                                        break;
                                }
                                catch (Exception)
                                {
                                }
                            }
                            if (!found)
                            {
                                // It's a namespace we are completing.
                                string ns = args[1] + "." + args[2];
                                foreach (Assembly asm in assemblies)
                                {
                                    try
                                    {
                                        foreach (Type t3 in asm.GetTypes())
                                        {
                                            if (t3.Namespace == null)
                                                continue;
                                            if (t3.Namespace.StartsWith(ns))
                                            {
                                                System.Console.WriteLine(ns);
                                                found = true;
                                                break;
                                            }
                                        }
                                        if (found)
                                            break;
                                    }
                                    catch (Exception)
                                    {}
                                }
                            }
                        }
                    }
                }
                catch (Exception e)
                {
                    reportError(e);
                }
                return true;
            }
        };
        private AppDomain ad = null;
        private string[] assemblies;
        private DateTime[] times = null;
        public MyAppDomain(string[] arg)
        {
            assemblies = arg;
            times = new DateTime[arg.Length];
            LoadAssemblies();
        }

        public void LoadAssemblies()
        {
            if (ad != null)
            {
                System.Console.Error.WriteLine("Unloading domain");
                AppDomain.Unload(ad);
            }

            ad = AppDomain.CreateDomain("MyAppDomain");
            object o = ad.CreateInstanceAndUnwrap(Assembly.GetExecutingAssembly().FullName, "CompleteSharp+MyAppDomain+Hack");
            Hack h = o as Hack;
            h.ad = this;
            int idx = 0;

            foreach (string a in assemblies)
            {
                try
                {
                    FileInfo fi = new FileInfo(a);
                    times[idx] = fi.LastWriteTime;
                    idx++;

                    System.Console.Error.WriteLine("Loading: " + a);
                    h.Load(a);
                }
                catch (Exception e)
                {
                    reportError(e);
                }
            }
        }

        public bool Execute(string[] args, ArrayList modules)
        {
            CheckUpdate();
            object o = ad.CreateInstanceAndUnwrap(Assembly.GetExecutingAssembly().FullName, "CompleteSharp+MyAppDomain+Hack");
            Hack h = o as Hack;
            h.ad = this;
            return h.Execute(args, modules);
        }

        bool forceReload = false;

        private void CheckUpdate()
        {
            // Yes, polling like this is a bit messy, however FileSystemWatcher didn't
            // work for me on OS X, there probably aren't that many assemblies to check
            // and the polling is only done after the user requests a completion, so
            // it shouldn't be too bad
            bool reload = forceReload;
            for (int i = 0; i < times.Length; i++)
            {
                try
                {
                    FileInfo fi = new FileInfo(assemblies[i]);
                    if (fi.LastWriteTime > times[i])
                    {
                        System.Console.Error.WriteLine("changed: " + assemblies[i] + ", " + fi.LastWriteTime + ", " + times[i]);
                        reload = true;
                        break;
                    }
                }
                catch (Exception e)
                {
                    reportError(e);
                }
            }
            if (reload)
            {
                forceReload = false;
                LoadAssemblies();
            }
        }
    }


    static void reportError(Exception e)
    {
        System.Console.Error.WriteLine("Exception: " + e.Message);
        System.Console.Error.WriteLine(e.StackTrace);
        System.Console.Error.WriteLine(sep);
    }

    public static void Main(string[] arg)
    {
        try
        {
            bool first = true;

            string[] argv = null;
            if (arg.Length > 0)
            {
                argv = arg[0].Split(new string[] {sep},  StringSplitOptions.None);
            }
            else
            {
                argv = new string[0];
            }
            MyAppDomain ad = new MyAppDomain(argv);

            while (true)
            {
                try
                {
                    if (!first)
                        // Just to indicate that there's no more output from the command and we're ready for new input
                        System.Console.WriteLine(sep);
                    first = false;
                    string command = System.Console.ReadLine();
                    System.Console.Error.WriteLine("got: " + command);
                    if (command == null)
                        break;
                    string[] args = Regex.Split(command, sep);
                    ArrayList modules = null;
                    if (args[0] == "-findclass")
                    {
                        modules = new ArrayList();
                        string line = null;
                        try
                        {
                            while ((line = System.Console.ReadLine()) != null)
                            {
                                if (line == sep)
                                    break;
                                modules.Add(line);
                            }
                        }
                        catch (Exception)
                        {}
                    }

                    if (!ad.Execute(args, modules))
                    {
                        break;
                    }
                }
                catch (Exception e)
                {
                    reportError(e);
                }
            }
        }
        catch (Exception e)
        {
            reportError(e);
        }
    }
}

